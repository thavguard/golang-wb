package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestParseFields(t *testing.T) {
	cases := []struct {
		spec     string
		expected []int
		wantErr  bool
	}{
		{"1", []int{1}, false},
		{"1,3-5,2,5", []int{1, 2, 3, 4, 5}, false},
		{" 2 , 4-6 ", []int{2, 4, 5, 6}, false},
		{"3-3", []int{3}, false},
		{"0", nil, true},
		{"-", nil, true},
		{"2-1", nil, true},
	}
	for _, c := range cases {
		fs, err := parseFields(c.spec)
		if c.wantErr {
			if err == nil {
				t.Fatalf("expected error for %q", c.spec)
			}
			continue
		}
		if err != nil {
			t.Fatalf("unexpected error for %q: %v", c.spec, err)
		}
		for _, i := range c.expected {
			if _, ok := fs[i]; !ok {
				t.Fatalf("expected field %d present for %q", i, c.spec)
			}
		}
	}
}

func TestSelectFields(t *testing.T) {
	set := parseFieldSpecMust("1,3-4")
	line := "a\tb\tc\td"
	res, ok := selectFields(line, "\t", set, false)
	if !ok || res != "a\tc\td" {
		t.Fatalf("unexpected: %v %q", ok, res)
	}
	res, ok = selectFields("no_delim", "\t", set, true)
	if ok {
		t.Fatalf("expected skip for separatedOnly")
	}
}

func TestRun(t *testing.T) {
	input := strings.NewReader(join(
		"a\tb\tc\td",
		"1\t2\t3",
		"no_delim",
		"x\ty",
	))
	var out bytes.Buffer
	set := parseFieldSpecMust("2-3")
	if err := run(input, &out, set, "\t", true); err != nil {
		t.Fatalf("run error: %v", err)
	}
	got := strings.TrimRight(out.String(), "\n")
	want := "b\tc\n2\t3\ny"
	if got != want {
		t.Fatalf("\nGOT:\n%s\nWANT:\n%s", got, want)
	}
}
