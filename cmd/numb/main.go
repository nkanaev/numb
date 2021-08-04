package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/nkanaev/numb/pkg/parser"
	"github.com/nkanaev/numb/pkg/value"
	"github.com/nkanaev/numb/pkg/runtime"
	"golang.org/x/term"
	_ "embed"
)

var prompt = "> "
var prefix = "  "

var loadfiles = ""
var prec = 2

//go:embed builtin.txt
var builtin string

func repl(env map[string]value.Value) {
	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), state)

	screen := struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}
	terminal := term.NewTerminal(screen, prompt)
	terminal.Write([]byte("enter `q` to quit\n"))

	runtime := runtime.NewRuntime()
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
		out := runtime.Eval(line)
		terminal.Write([]byte(prefix + out + "\n"))
	}
}

func load(env map[string]value.Value, r io.Reader, name string) {
	scanner := bufio.NewScanner(r)
	linenum := 0
	for scanner.Scan() {
		linenum += 1
		line := scanner.Text()
		line = strings.SplitN(line, "|", 2)[0]
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		_, err := parser.Eval(line, env)
		if err != nil {
			log.Fatalf("load %s (line %d): %s", name, linenum, err)
		}
	}
}

func read(env map[string]value.Value, r io.Reader) {
	var qwidth, awidth int
	qlines := make([]string, 0)
	alines := make([]string, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.SplitN(line, "|", 2)[0]
		line = strings.TrimSpace(line)
		qlines = append(qlines, line)
		if len(line) > qwidth {
			qwidth = len(line)
		}
	}

	for _, line := range qlines {
		val, err := parser.Eval(line, env)
		if err == nil {
			out := ""
			if val.Fmt == value.DEC {
				out = val.Format(",", prec)
			} else {
				out = val.String()
			}
			alines = append(alines, out)

			outwidth := len(out)
			if val.Unit != nil {
				outwidth -= len(val.Unit.String())
			}

			if outwidth > awidth {
				awidth = outwidth
			}
		} else {
			alines = append(alines, "")
		}
	}

	for i := 0; i < len(qlines); i++ {
		q, a := qlines[i], alines[i]

		apad := awidth - len(a)
		if apad < 0 {
			apad = -apad - 1
		}
		if len(a) > 0 {
			fmt.Printf("%s%s    | %s%s\n",
				q, strings.Repeat(" ", qwidth-len(q)),
				strings.Repeat(" ", apad), a,
			)
		} else {
			fmt.Println(q)
		}
	}
}

func main() {
	flag.StringVar(&loadfiles, "load", os.Getenv("NUMB_LOAD"), "list of files to preload")
	flag.Parse()

	env := make(map[string]value.Value)

	load(env, strings.NewReader(builtin), "<builtin>")

	if loadfiles != "" {
		for _, path := range strings.Split(loadfiles, ";") {
			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			load(env, file, path)
		}
	}

	if flag.NArg() == 1 {
		path := flag.Arg(0)
		file, err := os.Open(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer file.Close()
		read(env, file)
		return
	}

	if term.IsTerminal(0) {
		repl(env)
	} else {
		read(env, os.Stdin)
	}
}
