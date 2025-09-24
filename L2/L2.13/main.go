package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type fieldSet map[int]struct{}

func parseFields(spec string) (fieldSet, error) {
	if strings.TrimSpace(spec) == "" {
		return nil, errors.New("empty field spec")
	}
	set := make(fieldSet)
	parts := strings.Split(spec, ",")
	for _, part := range parts {
		p := strings.TrimSpace(part)
		if p == "" {
			continue
		}
		if strings.Contains(p, "-") {
			bounds := strings.SplitN(p, "-", 2)
			startStr := strings.TrimSpace(bounds[0])
			endStr := strings.TrimSpace(bounds[1])
			if startStr == "" || endStr == "" {
				return nil, fmt.Errorf("invalid range: %q", p)
			}
			start, err1 := strconv.Atoi(startStr)
			end, err2 := strconv.Atoi(endStr)
			if err1 != nil || err2 != nil || start <= 0 || end <= 0 || end < start {
				return nil, fmt.Errorf("invalid range: %q", p)
			}
			for i := start; i <= end; i++ {
				set[i] = struct{}{}
			}
			continue
		}
		idx, err := strconv.Atoi(p)
		if err != nil || idx <= 0 {
			return nil, fmt.Errorf("invalid field: %q", p)
		}
		set[idx] = struct{}{}
	}
	return set, nil
}

func selectFields(line, delimiter string, set fieldSet, separatedOnly bool) (string, bool) {
	if separatedOnly && !strings.Contains(line, delimiter) {
		return "", false
	}
	fields := strings.Split(line, delimiter)
	var out []string
	for i := 1; i <= len(fields); i++ {
		if _, ok := set[i]; ok {
			out = append(out, fields[i-1])
		}
	}
	return strings.Join(out, delimiter), true
}

func run(r io.Reader, w io.Writer, set fieldSet, delimiter string, separatedOnly bool) error {
	if delimiter == "" {
		return errors.New("empty delimiter")
	}
	scanner := bufio.NewScanner(r)
	buf := make([]byte, 0, 1024*1024)
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {
		line := scanner.Text()
		res, ok := selectFields(line, delimiter, set, separatedOnly)
		if ok {
			fmt.Fprintln(w, res)
		}
	}
	return scanner.Err()
}

func main() {
	var fSpec string
	var delimiter string
	var separatedOnly bool

	flag.StringVar(&fSpec, "f", "", "fields to select (e.g. 1,3-5)")
	flag.StringVar(&delimiter, "d", "\t", "field delimiter (default: tab)")
	flag.BoolVar(&separatedOnly, "s", false, "only lines containing the delimiter")
	flag.Parse()

	if fSpec == "" {
		fmt.Fprintln(os.Stderr, "-f is required")
		os.Exit(2)
	}

	set, err := parseFields(fSpec)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	if delimiter == "\\t" {
		delimiter = "\t"
	}
	if len(delimiter) != 1 {
		fmt.Fprintln(os.Stderr, "delimiter must be a single character")
		os.Exit(2)
	}

	if err := run(os.Stdin, os.Stdout, set, delimiter, separatedOnly); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
