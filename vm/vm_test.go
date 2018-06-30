package vm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/tesujiro/exp_yacc/parser"
)

func main() {
	fmt.Println("vim-go")
}

func TestNumbers(t *testing.T) {
	tests := []struct {
		script     string
		result     interface{}
		errMessage string
	}{
		{script: "1", result: int64(1)},
		{script: "1.1", result: float64(1.1)},
		{script: "123", result: int64(123)},
		{script: "123.456", result: float64(123.456)},
		{script: "\"abc\"", result: "abc"},
		{script: "+1+4", result: int64(5)},
		{script: "-1+3", result: int64(2)},
		{script: "1+1", result: int64(2)},
		{script: "1+1.1", result: float64(2.1)},
		{script: "1.1+4", result: float64(5.1)},
		{script: "1.1+1.1", result: float64(2.2)},
		{script: "3-1.1", result: float64(1.9)},
		{script: "2.2-1.1", result: float64(1.1)},
		{script: "3-1-1", result: int64(1)},
		{script: "3-(1-1)", result: int64(3)},
		{script: "3*5", result: int64(15)},
		{script: "1.5*2", result: float64(3)},
		{script: "5*1.2", result: float64(6)},
		{script: "15/5", result: int64(3)},
		{script: "16/5", result: int64(3)},
		{script: "3/1.5", result: float64(2)},
		{script: "3/0", errMessage: "devision by zero"},
		{script: "15%5", result: int64(0)},
		{script: "16%5", result: int64(1)},
		{script: "15%4.1", result: int64(3)},
		{script: "\"a b c\"+\" d e f\"", result: "a b c d e f"},
		{script: "\"a b c\"-\" d e f\"", result: float64(0)},
		{script: "15.2%7.1", result: int64(1)},
		{script: "a=1;b=2;a+b", result: int64(3)},
		{script: "a=1;b=2;a+1==b", result: true},
		{script: "a=1;b=2;a+1!=b", result: false},
		{script: "a=1;b=2;a<b", result: true},
		{script: "a=1;b=1;a<=b", result: true},
		{script: "a=1;b=2;a>b", result: false},
		{script: "a=1;b=1;a>=b", result: true},
		{script: "a=1;b=0.1;a>b", result: true},
		{script: "a=1;b=0.1;c=15;(a+b)*c", result: float64(16.5)},
		{script: "a=1;b=0.1;c=15;(a+b)*c/0.5", result: float64(33)},
		{script: "a=1;if a==1 { a=2 }", result: int64(2)},
		{script: "a=1;if a==1 { a=2 };a", result: int64(2)},
		{script: "a=1;if a==1 { env_test=2 };env_test", errMessage: "unknown symbol 'env_test'"},
		{script: "a=2;if a==1 { a=2 } else { a=3;b=4 }", result: int64(4)},
		{script: "a=1;b=1;if a==1 { b=2 };b", result: int64(2)},
		{script: "a=1;if a==1 { a=11 } else if a==2 { a=12 } else if a==3 { a=13 };a", result: int64(11)},
		{script: "a=2;if a==1 { a=11 } else if a==2 { a=12 } else if a==3 { a=13 };a", result: int64(12)},
		{script: "a=3;if a==1 { a=11 } else if a==2 { a=12 } else if a==3 { a=13 };a", result: int64(13)},
		{script: "a=1;if a==1 { env_test=11 } else if a==2 { env_test=12 } else { env_test=13 };env_test", errMessage: "unknown symbol 'env_test'"},
		{script: "a=2;if a==1 { env_test=11 } else if a==2 { env_test=12 } else { env_test=13 };env_test", errMessage: "unknown symbol 'env_test'"},
		{script: "a=3;if a==1 { env_test=11 } else if a==2 { env_test=12 } else { env_test=13 };env_test", errMessage: "unknown symbol 'env_test'"},
	}
	for _, test := range tests {
		env := NewEnv()
		l := new(parser.Lexer)
		l.Init(strings.NewReader(test.script))
		parser.Parse(l)
		if actual, err := Run(l.Result, env); err != nil {
			if test.errMessage == "" || err.Error() != test.errMessage {
				t.Errorf("Run error:%#v want%#v - script:%v\n", err, test.errMessage, test.script)
			}
		} else if actual != test.result {
			t.Errorf("got %#v\nwant %#v - script: %v", actual, test.result, test.script)
			continue
		}
	}
}
