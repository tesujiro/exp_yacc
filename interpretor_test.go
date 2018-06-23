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
		{stdin: "1+1", stdout: "2"},
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
