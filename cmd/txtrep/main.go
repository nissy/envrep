package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/nissy/txtrep"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	version   string
	isVersion = flag.Bool("v", false, "show version and exit")
	isHelp    = flag.Bool("h", false, "this help")
)

func main() {
	os.Exit(exitcode(run()))
}

func exitcode(err error) int {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return 1
	}

	return 0
}

func run() error {
	flag.Parse()
	args := flag.Args()

	if *isVersion {
		fmt.Println("v" + version)
		return nil
	}

	if *isHelp {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] textfile\n", os.Args[0])
		flag.PrintDefaults()
		return nil
	}

	file := os.Stdin

	if len(args) > 0 {
		f, err := os.Open(args[0])

		if err != nil {
			return err
		}

		file = f
	} else {
		if terminal.IsTerminal(0) {
			return nil
		}
	}

	defer file.Close()

	stat, err := file.Stat()

	if err != nil {
		return err
	}

	b := make([]byte, 0, stat.Size())
	r := bufio.NewReader(file)

	for {
		line, err := r.ReadBytes('\n')

		if err != nil && err != io.EOF {
			return err
		}

		b = append(b, line...)

		if err == io.EOF {
			break
		}
	}

	fmt.Print(txtrep.Replace(string(b)))

	return nil
}
