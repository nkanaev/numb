package main

import (
	"errors"
	"flag"
	"io"
	"os"
	"strings"

	"github.com/nkanaev/numb/pkg/parser"
	"github.com/nkanaev/numb/pkg/value"
	"golang.org/x/term"
)

var prompt = "> "
var prefix = "  "

var sep = ","
var prec = 2

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

	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
			panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), state)

	screen := struct{
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}
	terminal := term.NewTerminal(screen, prompt)
	terminal.Write([]byte("enter `q` to quit\n"))
	for {
		line, err := terminal.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			os.Stderr.WriteString(line + "\n")
			os.Exit(1)
		}
		line = strings.TrimSpace(line)
		if line == "q" {
			break
		}
		if line == "" {
			continue
		}
		val, err := eval(line, env)
		out := ""
		if err != nil {
			out = err.Error()
		} else {
			if val.Fmt == value.DEC {
				out = val.Format(sep, prec)
			} else {
				out = val.String()
			}
		}
		terminal.Write([]byte(prefix + out + "\n"))
	}
}

func main() {
	flag.IntVar(&prec, "prec", prec, "decimal precision")
	flag.StringVar(&sep, "sep", sep, "thousand separator")
	flag.Parse()
	repl()
}
