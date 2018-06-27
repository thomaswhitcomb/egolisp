package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"io"
	"os"
)

func evalSequence(env *envirs, cells *cell) *cell {
	var curr, head, ss *cell
	for s := cells; s != nil; s = s.next {
		ss = eval(env, s)
		if head == nil {
			head = ss
		} else {
			curr.next = ss
		}
		curr = ss
	}
	return head
}
func evalList(env *envirs, s *cell) *cell {
	cells := s.list()
	// functor is a scalar
	fs := eval(env, cells)
	if fs.isLambda() {
		fn := fs.lambda()
		head := evalSequence(env, cells.next)
		return fn(env, head)
	}
	if fs.isSymbol() {
		var ok bool
		var fd *fundef
		if fd, ok = funcs[*fs.symbol()]; !ok {
			fmt.Printf("Can not find function: %s\n", *fs.symbol())
			return makeFalse()
		}
		if fd.arityMin > cells.next.size() || fd.arityMax < cells.next.size() {
			fmt.Printf("Invalid arity on function %s\n", *fs.symbol())
			return makeFalse()
		}
		if _, ok = specialFuncs[*fs.symbol()]; ok {
			return fd.ep(env, cells.next)
		} else {
			head := evalSequence(env, cells.next)
			return fd.ep(env, head)
		}
	}
	fmt.Printf("Can not find function: %s\n", *fs.show())
	return makeFalse()
}

func evalScalar(env *envirs, s *cell) *cell {
	var v *cell
	var b bool
	v = env.find(*s.show())
	if v != nil {
		return v.clone()
	}
	v, b = vars[*s.show()]
	if b {
		return v.clone()
	}
	return s.clone()
}

func eval(env *envirs, s *cell) *cell {
	if s == nil {
		return makeList(nil)
	}
	if s.isScalar() {
		return evalScalar(env, s)
	}
	return evalList(env, s)
}

func run(text string) []string {
	var ans []string = []string{}
	tokens := tokenize(text)
	aparser := newParser()
	aparser.setTokens(tokens)
	cells := aparser.parse()
	for cell := cells; cell != nil; cell = cell.next {
		env := newEnvirs(nil)
		cell1 := eval(env, cell)
		ans = append(ans, fmt.Sprintf("%v", *cell1.show()))
	}
	return ans
}

func cli(reader io.ByteReader) {

	var p int = 0
	var buf bytes.Buffer
	var prompt string = "> "
	fmt.Printf("%s", prompt)
	for {
		r, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("error io: %v\n", err)
			os.Exit(1)
		}
		if r == '(' {
			p = p + 1
		}
		if r == ')' {
			p = p - 1
		}
		if r == 10 {
			if p == 0 {
				results := run(buf.String())
				for _, line := range results {
					fmt.Printf("%s\n", line)
				}
				buf.Reset()
				fmt.Printf("%s", prompt)
			} else {
				buf.WriteByte(' ')
			}
		} else {
			buf.WriteByte(r)
		}
	}
}

type Request struct {
	Line string `json:"line"`
}

type Response struct {
	Lines  []string `json:"lines"`
	Status int      `json:"status"`
}

func Handler(request Request) (Response, error) {
	results := run(request.Line)
	return Response{
		Lines:  results,
		Status: 200,
	}, nil
}

func main() {
	if os.Getenv("LAMBDA_TASK_ROOT") == "" {
		if len(os.Args) == 2 {
			f, err := os.Open(os.Args[1])
			if err != nil {
				fmt.Printf("Whoops, IO error on that file\n")
				os.Exit(1)
			}
			cli(bufio.NewReader(f))
		}
		cli(bufio.NewReader(os.Stdin))
	} else {
		lambda.Start(Handler)
	}
}
