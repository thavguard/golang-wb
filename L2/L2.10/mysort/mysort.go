package mysort

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

// Опции

type mySort struct {
	// k — 1-based индекс колонки; 0 — без выбора колонки (сравниваем всю строку)
	k int
	n bool
	r bool
	u bool
}

type MySortOption func(*mySort)

func WithK(k int) MySortOption { return func(s *mySort) { s.k = k } }
func WithN() MySortOption      { return func(s *mySort) { s.n = true } }
func WithR() MySortOption      { return func(s *mySort) { s.r = true } }
func WithU() MySortOption      { return func(s *mySort) { s.u = true } }

// Конструктор

func NewMySort(opts ...MySortOption) *mySort {
	s := &mySort{
		k: 0,
		n: false,
		r: false,
		u: false,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// helpers

func getKey(line string, k int) string {
	if k <= 0 {
		return line
	}
	// Разделитель — табуляция
	cols := strings.Split(line, "\t")
	idx := k - 1 // 1-based -> 0-based
	if idx < 0 || idx >= len(cols) {
		return ""
	}
	return cols[idx]
}

func parseNum(s string) (float64, bool) {
	if s == "" {
		return math.NaN(), false
	}
	// Пытаемся как float (покроет int тоже)
	v, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return math.NaN(), false
	}
	return v, true
}

// Sort — применяет выбранные опции
func (s *mySort) Sort(in []string) []string {
	if len(in) <= 1 {
		if s.u {
			return uniqueStable(in)
		}
		return in
	}

	out := slices.Clone(in)

	slices.SortFunc(out, func(a, b string) int {
		ka := getKey(a, s.k)
		kb := getKey(b, s.k)

		var cmp int
		if s.n {
			va, oka := parseNum(ka)
			vb, okb := parseNum(kb)

			switch {
			case !oka && okb:
				cmp = 1 // NaN после чисел
			case oka && !okb:
				cmp = -1
			case !oka && !okb:
				// оба не числа — лексикографически
				if ka < kb {
					cmp = -1
				} else if ka > kb {
					cmp = 1
				} else {
					cmp = 0
				}
			default:
				if va < vb {
					cmp = -1
				} else if va > vb {
					cmp = 1
				} else {
					cmp = 0
				}
			}
		} else {
			if ka < kb {
				cmp = -1
			} else if ka > kb {
				cmp = 1
			} else {
				cmp = 0
			}
		}

		if s.r {
			cmp = -cmp
		}
		return cmp
	})

	if s.u {
		out = uniqueStable(out)
	}

	return out
}

func uniqueStable(in []string) []string {
	seen := make(map[string]struct{}, len(in))
	out := make([]string, 0, len(in))
	for _, v := range in {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}
	return out
}
