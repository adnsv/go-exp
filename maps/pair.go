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

// SortStableByKey stable sorts key-value pairs by key.
func SortStableByKey[K constraints.Ordered, V any](pairs []*Pair[K, V]) {
	slices.SortStableFunc(pairs, func(a, b *Pair[K, V]) bool {
		return a.Key < b.Key
	})
}

// SortStableByKey stable sorts key-value pairs by key as determined by the less function.
func SortStableByKeyFunc[K any, V any](pairs []*Pair[K, V], less func(a, b K) bool) {
	slices.SortStableFunc(pairs, func(a, b *Pair[K, V]) bool {
		return less(a.Key, b.Key)
	})
}

// SortStableByKey stable sorts key-value pairs by value.
func SortStableByVal[K comparable, V constraints.Ordered](pairs []*Pair[K, V]) {
	slices.SortStableFunc(pairs, func(a, b *Pair[K, V]) bool {
		return a.Val < b.Val
	})
}

// SortStableByKey stable sorts key-value pairs by value as determined by the less function.
func SortStableByValFunc[K comparable, V any](pairs []*Pair[K, V], less func(a, b V) bool) {
	slices.SortStableFunc(pairs, func(a, b *Pair[K, V]) bool {
		return less(a.Val, b.Val)
	})
}

// Pairs returns a slice of key-value pairs constructed from m.
func Pairs[M ~map[K]V, K comparable, V any](m M) []*Pair[K, V] {
	pairs := make([]*Pair[K, V], 0, len(m))

	for k, v := range m {
		pairs = append(pairs, &Pair[K, V]{k, v})
	}

	return pairs
}

// StableSortedByKey returns a slice of key-value pairs constructed from m and
// stable-sorted by key.
func StableSortedByKey[M ~map[K]V, K constraints.Ordered, V any](m M) []*Pair[K, V] {
	ret := Pairs(m)
	SortStableByKey(ret)
	return ret
}

// StableSortedByKeyFunc returns a slice of key-value pairs constructed from m and
// stable-sorted by key as determined by the less function.
func StableSortedByKeyFunc[M ~map[K]V, K comparable, V any](m M, less func(a, b K) bool) []*Pair[K, V] {
	pairs := Pairs(m)
	SortStableByKeyFunc(pairs, less)
	return pairs
}

// StableSortedByVal returns a slice of key-value pairs constructed from m and
// stable-sorted by value.
func StableSortedByVal[M ~map[K]V, K constraints.Ordered, V constraints.Ordered](m M) []*Pair[K, V] {
	ret := Pairs(m)
	SortStableByVal(ret)
	return ret
}

// StableSortedByValFunc returns a slice of key-value pairs constructed from m and
// stable-sorted by value as determined by the less function.
func StableSortedByValFunc[M ~map[K]V, K constraints.Ordered, V any](m M, less func(a, b V) bool) []*Pair[K, V] {
	ret := Pairs(m)
	SortStableByValFunc(ret, less)
	return ret
}
