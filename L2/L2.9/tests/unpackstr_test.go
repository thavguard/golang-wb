package unpackstr_test

import (
	"testing"

	"L2.9/unpackstr"
)

func TestUnpack(t *testing.T) {

	type Data struct {
		input  string
		expect string
	}

	datas := []Data{{input: "a4bc2d5e", expect: "aaaabccddddde"}, {input: "abcd", expect: "abcd"}, {input: "45", expect: ""}, {input: "", expect: ""}, {input: `qwe\4\5`, expect: "qwe45"}, {input: `qwe\45`, expect: "qwe44444"}}

	for _, d := range datas {
		result, _ := unpackstr.Unpack(d.input)

		if result != d.expect {
			t.Errorf("SOME ERROR! EXPECT: %v; GOT: %v", d.expect, result)

		}

	}
}
