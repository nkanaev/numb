package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nkanaev/numb/pkg/parser"
	"github.com/nkanaev/numb/pkg/value"
)

var prompt = "> "

func eval(expr string, env map[string]value.Value) (value.Value, error) {
	tree := parser.Parse(expr)
	return tree.Eval(env), nil
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
		if err == nil {
			fmt.Println(" ", val)
		}
		fmt.Print(prompt)
	}
}

func main() {
	repl()
}
