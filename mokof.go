package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/ktr0731/go-fuzzyfinder"
)

func main() {
	var (
		bytearr bool
	)
	flag.BoolVar(&bytearr, "bytearr", false, "print as byte array")
	flag.Parse()

	os.Exit(run(bytearr))
}

func run(bytearr bool) int {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return 1
	}

	lines := fromPipe()
	if len(lines) == 0 {
		return 1
	}
	idx, err := fuzzyfinder.Find(lines, func(i int) string {
		return lines[i]
	})
	if err != nil {
		return 1
	}
	l := lines[idx]
	if bytearr {
		printBytes(l)
	} else {
		fmt.Println(l)
	}
	return 0
}

func printBytes(s string) {
	for _, b := range []byte(s) {
		fmt.Println(b)
	}
}

func fromPipe() []string {
	ss := []string{}
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		ss = append(ss, sc.Text())
	}
	if sc.Err() != nil {
		fmt.Println(sc.Err())
	}
	return ss
}
