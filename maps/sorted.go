package maps

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

// SortedKeys returns sorted keys of the map m. This sort is guaranteed to be
// stable because the keys are unique.
func SortedKeys[M ~map[K]V, K constraints.Ordered, V any](m M) []K {
	keys := maps.Keys(m)
	slices.Sort(keys)
	return keys
}

// SortedKeysFunc returns sorted keys of the map m as determined by the less
// function. This sort is stable provided the less function produces stable
// results.
func SortedKeysFunc[M ~map[K]V, K comparable, V any](m M, less func(a, b K) bool) []K {
	keys := maps.Keys(m)
	slices.SortFunc(keys, less)
	return keys
}
