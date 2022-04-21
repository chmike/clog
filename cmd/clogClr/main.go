package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/alexflint/go-arg"
)

var args struct {
	FileName string `arg:"positional" help:"clog file to display with colors"`
}

const (
	Reset      = "\033[0m"
	Bright     = "\033[1m"
	Dim        = "\033[2m"
	Underscore = "\033[4m"
	Blink      = "\033[5m"
	Reverse    = "\033[7m"
	Hidden     = "\033[8m"

	FgBlack   = "\033[30m"
	FgRed     = "\033[31m"
	FgGreen   = "\033[32m"
	FgYellow  = "\033[33m"
	FgBlue    = "\033[34m"
	FgMagenta = "\033[35m"
	FgCyan    = "\033[36m"
	FgWhite   = "\033[37m"

	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"
)

func main() {
	arg.MustParse(&args)
	var f *os.File
	if args.FileName == "" {
		f = os.Stdin
	} else {
		var err error
		f, err = os.Open(args.FileName)
		if err != nil {
			log.Fatal(err)
		}
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			fmt.Println(line)
			continue
		}
		switch line[0] {
		case 'F':
			fmt.Print(BgRed, FgBlack, line, Reset, "\n")
		case 'E':
			fmt.Print(Bright, FgRed, line, Reset, "\n")
		case 'W':
			fmt.Print(FgYellow, line, Reset, "\n")
		case 'I':
			fmt.Print(Bright, FgGreen, line, Reset, "\n")
		case ' ':
			fmt.Print(FgWhite, line, Reset, "\n")
		case '.', '-', '=':
			fmt.Print(Dim, FgWhite, line, Reset, "\n")
		default:
			fmt.Println(line)
		}
	}
}
