package main

import (
	"L2.12/conparams"
	"L2.12/mygrep"
)

func main() {
	params := conparams.NewParams()
	grep := mygrep.NewMygrep(params)

	grep.ReadFile()
}
