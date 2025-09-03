package mysort

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

type MyReader struct {
	Scanner *bufio.Scanner
	File    *os.File
}

func (r *MyReader) Scan() bool {
	return r.Scanner.Scan()
}

func (r *MyReader) Text() string {
	return r.Scanner.Text()
}

func NewMyReader() *MyReader {
	file, err := os.Open(filepath.Join("assets", "strings_50mb.tsv"))

	if err != nil {
		log.Fatal(err)
		return &MyReader{}
	}

	// defer file.Close()

	scanner := bufio.NewScanner(file)

	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	return &MyReader{Scanner: scanner, File: file}
}
