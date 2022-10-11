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
	return
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
	return
}

// CalcMerge calculates statistics for merging src into dst.
//
//   - Data in both dst and src remains unchanged
//   - Statistics is returned as sets of keys `map[K]struct{}`
//   - Keys in src that are not in dst are returned in the `create` set
//   - Keys/value pairs that are equal in both src and dst are returned in the `overwrite` set
//   - Keys that have different values in src and dst are returned in the `conflicts` set
//
func CalcMerge[M1 ~map[K]V, M2 ~map[K]V, K comparable, V comparable](dst M1, src M2) (create, overwrite, conflicts map[K]struct{}) {
	create = map[K]struct{}{}
	overwrite = map[K]struct{}{}
	conflicts = map[K]struct{}{}
	for k, v := range src {
		prev_v, exists := dst[k]
		if !exists {
			create[k] = struct{}{}
		} else if prev_v == v {
			overwrite[k] = struct{}{}
		} else {
			conflicts[k] = struct{}{}
		}
	}
	return
}

// CalcMergeFunc provides the same functionality as CalcMerge, but uses the
// allow functor to determine if a value can be overwritten. Notice also, that
// both dst and src maps in CalcMergeFunc are allowed to have different value
// types, which can help in merging heterogeneous data.
func CalcMergeFunc[M1 ~map[K]V1, M2 ~map[K]V2, K comparable, V1, V2 any](dst M1, src M2, allow func(key K, dstval V1, srcval V2) bool) (create, overwrite, conflicts map[K]struct{}) {
	create = map[K]struct{}{}
	overwrite = map[K]struct{}{}
	conflicts = map[K]struct{}{}
	for k, v := range src {
		prev_v, exists := dst[k]
		if !exists {
			create[k] = struct{}{}
		} else if allow(k, prev_v, v) {
			overwrite[k] = struct{}{}
		} else {
			conflicts[k] = struct{}{}
		}
	}
	return
}
