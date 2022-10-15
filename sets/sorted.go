package sets

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// Sorted returns the keys from s as a sorted slice. This sort is guaranteed to
// be stable.
func Sorted[S ~map[K]struct{}, K constraints.Ordered](s S) []K {
	r := Keys(s)
	slices.Sort(r)
	return r
}

// Sorted returns the keys from s as a sorted slice as determined by the less
// function. This sort is stable provided the less function produces stable
// results.
func SortedFunc[S ~map[K]struct{}, K constraints.Ordered](s S, less func(a, b K) bool) []K {
	r := Keys(s)
	slices.SortFunc(r, less)
	return r
}
