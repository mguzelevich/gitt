package parser

import (
	"bytes"
	"fmt"
	"strings"
)

func AnsiColor(name string, useAnsiColors bool) string {
	if !useAnsiColors {
		return ""
	}
	color := ""
	switch name {
	case "black":
		color = "30"
	case "red":
		color = "31"
	case "green":
		color = "32"
	case "yellow":
		color = "33"
	case "blue":
		color = "34"
	case "magenta":
		color = "35"
	case "cyan":
		color = "36"
	case "white":
		color = "37"
	}
	return color
}

func MakeFilesList(files []interface{}, prefix string) string {
	output := bytes.Buffer{}
	for _, f := range files {
		fmt.Fprintf(&output, "%s%s\n", prefix, f)
	}
	return output.String()
}

func StripMatches(matches []string) []string {
	result := []string{}
	for _, l := range matches {
		result = append(result, strings.TrimSpace(l))
	}
	return result
}
