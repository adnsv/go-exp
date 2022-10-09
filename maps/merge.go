package maps

// Merge copies key/value pairs in src adding them to dst. When a key from src
// is already in dst and the associated values are different, instead of
// overwriting, the whole key/value pair from src is copied to conflicts.
//
// Upon completion of this routine, the caller may analyze the conflicted keys
// and:
//
//   - Discard and reject the conflicts
//   - Overwrite the dst, for example by calling golang.org/x/exp/maps.Copy routine
//   - Implement more granular solution by merging each value individually
//
func Merge[M1 ~map[K]V, M2 ~map[K]V, K comparable, V comparable](dst M1, src M2) (conflicts M2) {
	conflicts = M2{}
	for k, v := range src {
		prev_v, exists := dst[k]
		if !exists || prev_v == v {
			dst[k] = v
		} else {
			conflicts[k] = v
		}
	}
	return conflicts
}

// MergeFunc provides the same functionality as Merge, but uses the allow
// functor to determine if a value can be overwritten.
func MergeFunc[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2, allow func(dstval, srcval V) bool) (conflicts M2) {
	conflicts = M2{}
	for k, v := range src {
		prev_v, exists := dst[k]
		if !exists || allow(prev_v, v) {
			dst[k] = v
		} else {
			conflicts[k] = v
		}
	}
	return conflicts
}
