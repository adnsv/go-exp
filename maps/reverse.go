package maps

func HasDuplicates[M ~map[K]V, K comparable, V comparable](m M) bool {
	return len(m) == len(valueSet(m))
}

func Inverted[M ~map[K]V, K comparable, V comparable](m M) (inv map[V]K, dup map[V]struct{}) {
	inv = map[V]K{}
	dup = map[V]struct{}{}
	for k, v := range m {
		_, exists := inv[v]
		if !exists {
			inv[v] = k
		} else {
			dup[v] = struct{}{}
		}
	}
	for k := range dup {
		delete(inv, k)
	}
	return
}

func KeysForVal[M ~map[K]V, K comparable, V comparable](m M, v V) map[K]struct{} {
	r := map[K]struct{}{}
	for k, val := range m {
		if val == v {
			r[k] = struct{}{}
		}
	}
	return r
}

func valueSet[M ~map[K]V, K comparable, V comparable](m M) map[V]struct{} {
	r := map[V]struct{}{}
	for _, v := range m {
		if _, exists := r[v]; !exists {
			r[v] = struct{}{}
		}
	}
	return r
}
