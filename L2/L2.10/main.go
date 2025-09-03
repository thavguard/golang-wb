package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"L2.10/mysort"
)

// expandShortFlags превращает "-ru" -> "-r" "-u" для совместимости с привычным синтаксисом.
func expandShortFlags(args []string) []string {
	out := make([]string, 0, len(args))
	out = append(out, args[0])
	for _, a := range args[1:] {
		if len(a) > 2 && a[0] == '-' && a[1] != '-' { // короткие склеенные флаги
			// Не трогаем флаги с параметром вида "-k5" — поддержим только "-k 5".
			for i := 1; i < len(a); i++ {
				out = append(out, "-"+string(a[i]))
			}
		} else {
			out = append(out, a)
		}
	}
	return out
}

func main() {
	// Нормализуем короткие склеенные флаги
	os.Args = expandShortFlags(os.Args)

	var (
		kFlag int
		nFlag bool
		rFlag bool
		uFlag bool
	)

	flag.IntVar(&kFlag, "k", 0, "сортировать по колонке №N (1-based), разделитель — табуляция")
	flag.BoolVar(&nFlag, "n", false, "числовая сортировка")
	flag.BoolVar(&rFlag, "r", false, "обратный порядок")
	flag.BoolVar(&uFlag, "u", false, "уникальные строки")
	flag.Parse()

	// Собираем опции сортировщика по мере использования флагов
	opts := make([]mysort.MySortOption, 0, 4)
	if kFlag > 0 {
		opts = append(opts, mysort.WithK(kFlag))
	}
	if nFlag {
		opts = append(opts, mysort.WithN())
	}
	if rFlag {
		opts = append(opts, mysort.WithR())
	}
	if uFlag {
		opts = append(opts, mysort.WithU())
	}

	r := mysort.NewMyReader()
	defer r.File.Close()

	sorter := mysort.NewMySort(opts...)

	var lines []string
	for r.Scan() {
		lines = append(lines, r.Text())
	}
	if err := r.Scanner.Err(); err != nil {
		log.Fatalf("read error: %v", err)
	}

	fmt.Printf("Прочитано строк: %d\n", len(lines))

	sorted := sorter.Sort(lines)

	outDir := "assets"
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		log.Fatalf("mkdir: %v", err)
	}
	outPath := filepath.Join(outDir, "result.txt")

	abs, _ := filepath.Abs(outPath)
	fmt.Printf("Пишу в файл: %s\n", abs)

	file, err := os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		log.Fatalf("open for write: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range sorted {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			log.Fatalf("write: %v", err)
		}
	}
	if err := writer.Flush(); err != nil {
		log.Fatalf("flush: %v", err)
	}
	if err := file.Sync(); err != nil {
		log.Fatalf("sync: %v", err)
	}

	fmt.Println("Готово.")
}
