package maps

// KeySet returns a set constructed from all the keys in m.
func KeySet[M ~map[K]V, K comparable, V any](m M) map[K]struct{} {
	r := make(map[K]struct{}, len(m))
	for k := range m {
		r[k] = struct{}{}
	}
	return r
}

// Sliced returns elements from m that exist in s.
func Sliced[M ~map[K]V, S ~map[K]struct{}, K comparable, V any](m M, s S) M {
	r := M{}
	for k := range s {
		if v, ok := m[k]; ok {
			r[k] = v
		}
	}
	return r
}
