package maps

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// Pair is a key-value pair that can be used to flatten maps.
type Pair[K any, V any] struct {
	Key K
	Val V
}

// Pairs returns a slice of key-value pairs constructed from m. The pairs will
// be in an indeterminate order.
func Pairs[M ~map[K]V, K comparable, V any](m M) []*Pair[K, V] {
	pairs := make([]*Pair[K, V], 0, len(m))

	for k, v := range m {
		pairs = append(pairs, &Pair[K, V]{k, v})
	}

	return pairs
}

// SortedFunc returns a slice of key-value pairs constructed from m and sorted
// as determined by the less function. The sort is stable if the less function
// produces stable results.
func SortedFunc[M ~map[K]V, K comparable, V any](m M, less func(a, b *Pair[K, V]) bool) []*Pair[K, V] {
	ret := Pairs(m)
	slices.SortFunc(ret, less)
	return ret
}

// SortedByKey returns a slice of key-value pairs constructed from m and sorted
// by key. This sort is guaranteed to be stable because the keys are unique.
func SortedByKey[M ~map[K]V, K constraints.Ordered, V any](m M) []*Pair[K, V] {
	return SortedFunc(m, lessByKey[K, V])
}

// SortedByKeyFunc returns a slice of key-value pairs constructed from m and
// sorted by key as determined by the less function. This sort is stable
// provided the less function produces stable results.
func SortedByKeyFunc[M ~map[K]V, K comparable, V any](m M, less func(a, b K) bool) []*Pair[K, V] {
	return SortedFunc(m, lessByKeyFunc[K, V](less))
}

// SortedByVal returns a slice of key-value pairs constructed from m and sorted
// by value. This sort is stable only if there is no duplicates (all values are
// unique).
func SortedByVal[M ~map[K]V, K comparable, V constraints.Ordered](m M) []*Pair[K, V] {
	ret := Pairs(m)
	slices.SortFunc(ret, lessByVal[K, V])
	return ret
}

// SortedByValFunc returns a slice of key-value pairs constructed from m and
// sorted by value as determined by the less function. This sort is stable
// provided there is no duplicates (all values are unique) and the less function
// produces stable results.
func SortedByValFunc[M ~map[K]V, K comparable, V any](m M, less func(a, b V) bool) []*Pair[K, V] {
	ret := Pairs(m)
	slices.SortFunc(ret, lessByValFunc[K](less))
	return ret
}

// StableSortedByVal returns a slice of key-value pairs constructed from m and
// sorted by value. For duplicate values, sorting falls back to comparing keys.
// This sort is guaranteed to be stable.
func StableSortedByVal[M ~map[K]V, K constraints.Ordered, V constraints.Ordered](m M) []*Pair[K, V] {
	ret := Pairs(m)
	slices.SortFunc(ret, lessByValKey[K, V])
	return ret
}

// SortedByValFunc returns a slice of key-value pairs constructed from m and
// sorted by value as determined by the less function. For duplicate values,
// sorting falls back to comparing keys. This sort is stable provided the less
// function produces stable results.
func StableSortedByValFunc[M ~map[K]V, K constraints.Ordered, V any](m M, less func(a, b V) bool) []*Pair[K, V] {
	ret := Pairs(m)
	slices.SortFunc(ret, lessByValKeyFunc[K](less))
	return ret
}

// Insert copies key-value pairs into m, if the map doesn't already contain an
// elements with equivalent keys.
func Insert[M ~map[K]V, K comparable, V any](m M, pairs ...*Pair[K, V]) {
	for _, p := range pairs {
		_, exists := m[p.Key]
		if !exists {
			m[p.Key] = p.Val
		}
	}
}

// InsertOrOverwrite copies key-value pairs into m, overwriting existing
// elements with equivalent keys.
func InsertOrOverwrite[M ~map[K]V, K comparable, V any](m M, pairs ...*Pair[K, V]) {
	for _, p := range pairs {
		m[p.Key] = p.Val
	}
}

// sorting callbacks, used internally

func lessByKey[K constraints.Ordered, V any](a, b *Pair[K, V]) bool {
	return a.Key < b.Key
}

func lessByVal[K any, V constraints.Ordered](a, b *Pair[K, V]) bool {
	return a.Val < b.Val
}

func lessByKeyFunc[K any, V any](less func(a, b K) bool) func(a, b *Pair[K, V]) bool {
	return func(a, b *Pair[K, V]) bool {
		return less(a.Key, b.Key)
	}
}

func lessByValFunc[K any, V any](less func(a, b V) bool) func(a, b *Pair[K, V]) bool {
	return func(a, b *Pair[K, V]) bool {
		return less(a.Val, b.Val)
	}
}

func lessByValKey[K constraints.Ordered, V constraints.Ordered](a, b *Pair[K, V]) bool {
	if a.Val < b.Val {
		return true
	} else if b.Val < a.Val {
		return false
	} else {
		return a.Key < b.Key
	}
}
func lessByValKeyFunc[K constraints.Ordered, V any](less func(a, b V) bool) func(a, b *Pair[K, V]) bool {
	return func(a, b *Pair[K, V]) bool {
		if less(a.Val, b.Val) {
			return true
		} else if less(b.Val, a.Val) {
			return false
		} else {
			return a.Key < b.Key
		}
	}
}
