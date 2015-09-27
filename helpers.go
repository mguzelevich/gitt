package main

import (
	"fmt"
	"os"
	"strings"
)

func printMatches(name string, matches []string) {
	fmt.Fprintf(os.Stderr, "> %s : %s\n", name, strings.Join(matches, "|"))
}
