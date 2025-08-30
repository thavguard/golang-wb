package myset

import "fmt"

type MySet struct {
	Set map[string]struct{}
}

func (myset *MySet) String() string {
	var result string

	result += "{"

	for k := range myset.Set {
		result += fmt.Sprintf(" %v,", k)
	}

	result += " }"

	return result
}

func (myset *MySet) Add(v string) bool {
	prevLen := len(myset.Set)

	myset.Set[v] = struct{}{}

	return prevLen != len(myset.Set)
}

func (myset *MySet) Append(v ...string) int {
	prevLen := len(myset.Set)

	for _, item := range v {
		myset.Set[item] = struct{}{}
	}

	return len(myset.Set) - prevLen
}

func (myset *MySet) Intersect(other map[string]struct{}) map[string]struct{} {
	intersect := make(map[string]struct{})

	if len(myset.Set) < len(other) {
		for k := range myset.Set {
			if _, ok := other[k]; ok {
				intersect[k] = struct{}{}
			}
		}
	} else {
		for k := range other {
			if _, ok := myset.Set[k]; ok {
				intersect[k] = struct{}{}
			}
		}
	}

	return intersect
}

func NewMySet() *MySet {
	return &MySet{Set: make(map[string]struct{})}
}
