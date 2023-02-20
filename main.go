package main

import (
	"flag"
	"fmt"
	"os"

	b64 "github.com/samharju/b64/internal"
)

var (
	decode     bool
	fileInput  string
	fileOutput string
)

var usage = `usage:
	b64 [-d] INPUTSTR
	b64 [-d] [-o <path>] INPUTSTR
	b64 [-d] [-i <path>]
	b64 [-d] [-i <path>] [-o <path>]

Reads from stdin and prints to stdout, unless input- or outputfile is given.
Encodes by default, decode with -d.
`

func main() {

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
		flag.PrintDefaults()
	}

	flag.BoolVar(&decode, "d", false, "decode input")
	flag.StringVar(&fileInput, "i", "", "read from file instead of stdin")
	flag.StringVar(&fileOutput, "o", "", "write to file instead of stdout")

	flag.Parse()

	var (
		err     error
		in, out []byte
	)

	if fileInput != "" {
		in, err = os.ReadFile(fileInput)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		if len(flag.Args()) != 1 {
			flag.Usage()
			os.Exit(1)
		}
		in = []byte(flag.Arg(0))
	}

	if decode {
		out, err = b64.DecodeBytes(in)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		out = b64.EncodeBytes(in)
	}

	if fileOutput != "" {
		err = os.WriteFile(fileOutput, out, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		fmt.Print(string(out))
	}
}
