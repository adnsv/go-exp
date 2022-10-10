package maps

// ValueSet returns a set constructed from all the values in m.
func ValueSet[M ~map[K]V, K comparable, V comparable](m M) map[V]struct{} {
	r := map[V]struct{}{}
	for _, v := range m {
		if _, exists := r[v]; !exists {
			r[v] = struct{}{}
		}
	}
	return r
}

// HasDuplicates checks whether m contains duplicates (multiple keys having the
// same value).
func HasDuplicates[M ~map[K]V, K comparable, V comparable](m M) bool {
	return len(m) == len(ValueSet(m))
}

// Inverted produces inverted map from m. Entries that can not be inverted are
// returned as a set of duplicates, they are excluded from the inverted result:
//
//   len(m) = len(inverted) + len(duplicates)
//
// A strategy for resolving the issues with duplicates then may include
// iterating over the returned set of duplicates, possibly calling the
// MatchValue function to discover which keys are associated to each duplicate
// and taking appropriate actions.
//
func Inverted[M ~map[K]V, K comparable, V comparable](m M) (inverted map[V]K, duplicates map[V]struct{}) {
	inverted = map[V]K{}
	duplicates = map[V]struct{}{}
	for k, v := range m {
		_, exists := inverted[v]
		if !exists {
			inverted[v] = k
		} else {
			duplicates[v] = struct{}{}
		}
	}
	for k := range duplicates {
		delete(inverted, k)
	}
	return
}

// MatchValue returns a set of keys that contain the same value v.
func MatchValue[M ~map[K]V, K comparable, V comparable](m M, v V) (keys map[K]struct{}) {
	keys = map[K]struct{}{}
	for k, val := range m {
		if val == v {
			keys[k] = struct{}{}
		}
	}
	return
}
