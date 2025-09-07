package mygrep

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"L2.12/conparams"
)

type mygrep struct {
	file   *os.File
	result []string
	params *conparams.Params
}

func (g mygrep) buildFlags() string {
	var flagsSb strings.Builder

	flagsSb.WriteString("(?")

	if g.params.I {
		flagsSb.WriteString("i")
	}

	flagsSb.WriteString(")")

	flags := flagsSb.String()

	// fmt.Printf("flags: %v\n", flags)

	return flags
}

func (g mygrep) buildPattern(flags string) string {
	var patternSb strings.Builder

	patternSb.WriteString(flags)

	if g.params.F {
		patternSb.WriteString("\\b")
	}

	patternSb.WriteString(g.params.Pattern)

	if g.params.F {
		patternSb.WriteString("\\b")
	}

	pattern := patternSb.String()

	// fmt.Printf("pattern: %v\n", pattern)

	return pattern
}

func (g *mygrep) checkStrings(lines []string) []string {

	// Build flags

	flags := g.buildFlags()

	// Build pattern

	pattern := g.buildPattern(flags)

	reg, err := regexp.Compile(pattern)

	if err != nil {
		log.Fatalf("Some error in regex: %v\n", err)
	}

	filter := g.params.V

	for index, line := range lines {

		if match := reg.MatchString(line); match == !filter {

			paramAfter := g.params.A
			paramBefore := g.params.B

			if g.params.C != 0 {
				paramC := (g.params.C + 1) / 2

				paramAfter = paramC
				paramBefore = paramC
			}

			if paramAfter != 0 {
				// Ensure we don't go below index 0
				start := max(index-paramAfter, 0)
				for iA := start; iA < index; iA++ {
					g.result = append(g.result, lines[iA])
				}
			}

			g.result = append(g.result, line)

			if paramBefore != 0 {
				// Ensure we don't go beyond the end of the slice
				end := index + paramBefore
				if end >= len(lines) {
					end = len(lines) - 1
				}
				for iA := index + 1; iA <= end; iA++ {
					g.result = append(g.result, lines[iA])
				}
			}

		}
	}

	return g.result
}

func (g *mygrep) ReadFile() {
	reader := bufio.NewReader(g.file)
	defer g.file.Close()

	content, err := io.ReadAll(reader)

	if err != nil {
		log.Fatalf("Something went wrong: %v\n", err)
	}

	lines := strings.Split(string(content), "\n")

	result := g.checkStrings(lines)

	for index, r := range result {

		var resultSb strings.Builder

		if g.params.N || g.params.Count {
			lineNumber := fmt.Sprintf("%d ", index)
			resultSb.WriteString(lineNumber)
		}

		if !g.params.Count {
			resultSb.WriteString("- " + r)
		}

		r = resultSb.String()

		fmt.Printf("%v\n", r)
	}

}

func NewMygrep(params *conparams.Params) *mygrep {
	file, err := os.Open(filepath.Join("assets", "file.txt"))

	if err != nil {
		log.Fatalf("Some error in open file: %v\n", err)
	}

	g := &mygrep{file: file, params: params}

	return g
}
