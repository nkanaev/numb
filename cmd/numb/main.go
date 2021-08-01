package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/nkanaev/numb/pkg/parser"
	"github.com/nkanaev/numb/pkg/unit"
	"github.com/nkanaev/numb/pkg/value"
	"golang.org/x/term"
)

var prompt = "> "
var prefix = "  "

var rates = ""
var sep = ","
var prec = 2

func repl() {
	env := make(map[string]value.Value)

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
		val, err := parser.Eval(line, env)
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

func read(r io.Reader) {
	env := make(map[string]value.Value)

	var qwidth, awidth int
	qlines := make([]string, 0)
	alines := make([]string, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.SplitN(line, "|", 2)[0]
		line = strings.TrimRightFunc(line, unicode.IsSpace)
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
				out = val.Format(sep, prec)
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

func loadCurrencies(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
	defer file.Close()

	currencies := make([]unit.Currency, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), "\t", " ")
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			continue
		}
		code := parts[0]
		rate, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			continue
		}
		currencies = append(currencies, unit.Currency{code, rate})
	}
	unit.AddExchangeRates(currencies)
}

func main() {
	var showUnits bool
	flag.IntVar(&prec, "prec", prec, "decimal precision")
	flag.StringVar(&sep, "sep", sep, "thousand separator")
	flag.StringVar(&rates, "rates", os.Getenv("NUMB_RATES"), "path to exchange rates file")
	flag.BoolVar(&showUnits, "units", false, "show available units and exit")
	flag.Parse()

	if showUnits {
		unit.Help()
		return
	}

	if rates != "" {
		loadCurrencies(rates)
	}

	if flag.NArg() == 1 {
		path := flag.Arg(0)
		file, err := os.Open(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer file.Close()
		read(file)
		return
	}

	if term.IsTerminal(0) {
		repl()
	} else {
		read(os.Stdin)
	}
}
