package myset

type MySet struct {
	Set map[int]struct{}
}

func (myset *MySet) Add(v int) bool {
	prevLen := len(myset.Set)

	myset.Set[v] = struct{}{}

	return prevLen != len(myset.Set)
}

func (myset *MySet) Append(v ...int) int {
	prevLen := len(myset.Set)

	for _, item := range v {
		myset.Set[item] = struct{}{}
	}

	return len(myset.Set) - prevLen
}

func (myset *MySet) Intersect(other map[int]struct{}) map[int]struct{} {
	intersect := make(map[int]struct{})

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
	return &MySet{Set: make(map[int]struct{})}
}
