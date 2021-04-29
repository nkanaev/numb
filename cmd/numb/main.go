package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/nkanaev/numb/pkg/parser"
	"github.com/nkanaev/numb/pkg/value"
)

var prompt = "> "

func eval(expr string, env map[string]value.Value) (val value.Value, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknown error")
			}
		}
	}()
	val = parser.Parse(expr).Eval(env)
	return
}

func repl() {
	env := make(map[string]value.Value)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("enter `q` to quit")
	fmt.Print(prompt)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "q" {
			break
		}
		val, err := eval(line, env)
		if err != nil {
			fmt.Println(" ", err)
		} else {
			fmt.Println(" ", val)
		}
		fmt.Print(prompt)
	}
}

func main() {
	repl()
}
