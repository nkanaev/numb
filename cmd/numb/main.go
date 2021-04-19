package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var prompt = "> "

type Value struct {
	num float64
}

func (v *Value) String() string {
	return strconv.FormatInt(int64(v.num), 10)
}

func eval(expr string, vars map[string]Value) *Value {
	return nil	
}

func repl() {
	vars := make(map[string]Value)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("enter `q` to quit")
	fmt.Print(prompt)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "q" {
			break
		}
		val := eval(line, vars)
		if val != nil {
			fmt.Println(" ", val.String())
		}
		fmt.Print(prompt)
	}
}

func main() {
	repl()
}
