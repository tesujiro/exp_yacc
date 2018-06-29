package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func TestInterpretor(t *testing.T) {
	cases := []struct {
		stdin  string
		stderr string
		stdout string
	}{
		{stdin: "+1+4", stdout: "5"},
		{stdin: "-1+3", stdout: "2"},
		{stdin: "1+1", stdout: "2"},
		{stdin: "1+1.1", stdout: "2.1"},
		{stdin: "1.1+4", stdout: "5.1"},
		{stdin: "1.1+1.1", stdout: "2.2"},
		{stdin: "3-1.1", stdout: "1.9"},
		{stdin: "2.2-1.1", stdout: "1.1"},
		{stdin: "3-1-1", stdout: "1"},
		{stdin: "3-(1-1)", stdout: "3"},
		{stdin: "3*5", stdout: "15"},
		{stdin: "1.5*2", stdout: "3"},
		{stdin: "5*1.2", stdout: "6"},
		{stdin: "15/5", stdout: "3"},
		{stdin: "16/5", stdout: "3"},
		{stdin: "3/1.5", stdout: "2"},
		{stdin: "3/0", stdout: "Eval error:devision by zero"},
		{stdin: "15%5", stdout: "0"},
		{stdin: "16%5", stdout: "1"},
		{stdin: "15%4.1", stdout: "3"},
		{stdin: "15.2%7.1", stdout: "1"},
		{stdin: "a=1;b=2;a+b", stdout: "3"},
		{stdin: "a=1;b=2;a+1==b", stdout: "true"},
		{stdin: "a=1;b=2;a+1!=b", stdout: "false"},
		{stdin: "a=1;b=2;a<b", stdout: "true"},
		{stdin: "a=1;b=1;a<=b", stdout: "true"},
		{stdin: "a=1;b=2;a>b", stdout: "false"},
		{stdin: "a=1;b=1;a>=b", stdout: "true"},
		{stdin: "a=1;b=0.1;a>b", stdout: "true"},
		{stdin: "a=1;b=0.1;c=15;(a+b)*c", stdout: "16.5"},
		{stdin: "a=1;b=0.1;c=15;(a+b)*c/0.5", stdout: "33"},
		{stdin: "a=1;if a==1 { a=2 }", stdout: "2"},
		{stdin: "a=1;if a==1 { a=2 };a", stdout: "2"},
		{stdin: "a=1;if a==1 { env_test=2 };env_test", stdout: "Eval error:unknown symbol 'env_test'"},
		{stdin: "a=2;if a==1 { a=2 } else { a=3;b=4 }", stdout: "4"},
		{stdin: "a=1;b=1;if a==1 { b=2 };b", stdout: "2"},
		{stdin: "a=1;if a==1 { a=11 } else if a==2 { a=12 } else if a==3 { a=13 };a", stdout: "11"},
		{stdin: "a=2;if a==1 { a=11 } else if a==2 { a=12 } else if a==3 { a=13 };a", stdout: "12"},
		{stdin: "a=3;if a==1 { a=11 } else if a==2 { a=12 } else if a==3 { a=13 };a", stdout: "13"},
		{stdin: "a=1;if a==1 { env_test=11 } else if a==2 { env_test=12 } else { env_test=13 };env_test", stdout: "Eval error:unknown symbol 'env_test'"},
		{stdin: "a=2;if a==1 { env_test=11 } else if a==2 { env_test=12 } else { env_test=13 };env_test", stdout: "Eval error:unknown symbol 'env_test'"},
		{stdin: "a=3;if a==1 { env_test=11 } else if a==2 { env_test=12 } else { env_test=13 };env_test", stdout: "Eval error:unknown symbol 'env_test'"},
	}
	realStdin := os.Stdin
	realStdout := os.Stdout
	realStderr := os.Stderr

	readFromIn, writeToIn, err := os.Pipe()
	if err != nil {
		t.Fatal("Pipe error:", err)
	}
	os.Stdin = readFromIn
	//logger.Print("pipe in created")

	readFromOut, writeToOut, err := os.Pipe()
	if err != nil {
		os.Stdin = realStdin
		os.Stderr = realStderr
		t.Fatal("Pipe error:", err)
	}
	os.Stdout = writeToOut
	//logger.Print("pipe out created")

	// out checker goroutine
	readerOut := bufio.NewReaderSize(readFromOut, 256)
	chanDone := make(chan struct{})
	chanOut := make(chan string)
	go func() {
		for {
			if data, err := readerOut.ReadString('\n'); err != nil && err != io.EOF {
				t.Fatal("ReadString error:", err)
			} else {
				select {
				case <-chanDone:
					return
				default:
				}
				chanOut <- data
			}
		}
	}()

	readTimeout := 10 * time.Millisecond
	time.Sleep(readTimeout)
	go run()

	for _, c := range cases {
		_, err := writeToIn.WriteString(c.stdin + "\n")
		if err != nil {
			os.Stdin = realStdin
			os.Stderr = realStderr
			os.Stdout = realStdout
			t.Fatal("Stdin WriteString error:", err)
		}

		select {
		case dataOut := <-chanOut:
			dataOut = strings.TrimSpace(dataOut)
			if dataOut != c.stdout {
				os.Stdin = realStdin
				os.Stderr = realStderr
				os.Stdout = realStdout
				t.Fatalf("Stdout - received: %v - expected: %v - runSource: %v", dataOut, c.stdout, c.stdin)
			}
		case <-time.After(readTimeout):
			if c.stdout != "" {
				os.Stdin = realStdin
				os.Stderr = realStderr
				os.Stdout = realStdout
				t.Fatalf("Stdout - received: %v - expected: %v - runSource: %v", "", c.stdout, c.stdin)
			}
		}
	}

	_, err = writeToIn.WriteString("exit\n")
	if err != nil {
		os.Stdin = realStdin
		os.Stderr = realStderr
		os.Stdout = realStdout
		t.Fatal("Stdin WriteString error:", err)
	}

	close(chanDone)
	writeToOut.Close()

	os.Stdin = realStdin
	os.Stderr = realStderr
	os.Stdout = realStdout
}
