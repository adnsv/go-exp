package sets

// set is effectively a map[K comparable]struct{}

// Contains checks if there is an element in the set.
func Contains[S ~map[K]struct{}, K comparable](s S, k K) bool {
	_, ok := s[k]
	return ok
}

// Contains checks if any of the elements is in the set.
func ContainsAny[S ~map[K]struct{}, K comparable](s S, keys ...K) bool {
	for _, k := range keys {
		_, ok := s[k]
		if ok {
			return true
		}
	}
	return false
}

// Contains checks if all the elements are in the set.
func ContainsAll[S ~map[K]struct{}, K comparable](s S, keys ...K) bool {
	for _, k := range keys {
		_, ok := s[k]
		if !ok {
			return false
		}
	}
	return true
}

// Equal reports whether two sets contain the same elements.
func Equal[S1, S2 ~map[K]struct{}, K comparable](s1 S1, s2 S2) bool {
	if len(s1) != len(s2) {
		return false
	}
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			return false
		}
	}
	return true
}

// Clear removes all elements from s.
func Clear[S ~map[K]struct{}, K comparable](s S) {
	for k := range s {
		delete(s, k)
	}
}

// Clone returns a copy of s.
func Clone[S ~map[K]struct{}, K comparable](s S) S {
	r := make(S, len(s))
	for k := range s {
		r[k] = struct{}{}
	}
	return r
}

// Elements returns the elements of the set s.
func Elements[S ~map[K]struct{}, K comparable](s S) []K {
	r := make([]K, 0, len(s))
	for k := range s {
		r = append(r, k)
	}
	return r
}

// Insert inserts the elements info the set.
func Insert[S ~map[K]struct{}, K comparable](s S, keys ...K) {
	for _, k := range keys {
		s[k] = struct{}{}
	}
}

// Remove removes the elements from the set.
func Remove[S ~map[K]struct{}, K comparable](s S, keys ...K) {
	for _, k := range keys {
		delete(s, k)
	}
}

// Union combines elements from s1 and s2 into one set.
// Returns s1 ∪ s2.
func Union[S map[K]struct{}, K comparable](s1 S, s2 S) S {
	r := Clone(s1)
	for k := range s2 {
		r[k] = struct{}{}
	}
	return r
}

// Merge inserts the src elements into the dst.
// Effectively, dst = dst ∪ src.
func Merge[S1 ~map[K]struct{}, S2 ~map[K]struct{}, K comparable](dst S1, src S2) {
	for k := range src {
		dst[k] = struct{}{}
	}
}

// Difference returns the elements from s1 that are not contained in s2.
// Returns s1 - s2.
func Difference[S map[K]struct{}, K comparable](s1 S, s2 S) S {
	r := S{}
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			r[k] = struct{}{}
		}
	}
	return r
}

// Subtract removes the src elements from the dst.
// Effectively, this is a difference (subtraction) operation: dst = dst - src.
func Subtract[S1 ~map[K]struct{}, S2 ~map[K]struct{}, K comparable](dst S1, src S2) {
	for k := range src {
		delete(dst, k)
	}
}

// Intersection returns elements that exist in both s1 and s2.
// Effectively: s1 ∩ s2
func Intersection[S map[K]struct{}, K comparable](s1 S, s2 S) S {
	n := len(s1)
	n2 := len(s2)
	if n == 0 || n2 == 0 {
		return S{}
	}
	if n2 < n {
		n = n2
	}
	r := make(S, n)
	for k := range s1 {
		if _, ok := s2[k]; ok {
			r[k] = struct{}{}
		}
	}
	return r
}

// Intersect removes elements from dst that are not contained in src.
// Effectively, dst = dst ∩ src
func Intersect[S1 ~map[K]struct{}, S2 ~map[K]struct{}, K comparable](dst S1, src S2) {
	if len(src) == 0 {
		Clear(dst)
	}
	if len(dst) == 0 {
		return
	}
	for k := range Difference(map[K]struct{}(dst), map[K]struct{}(src)) {
		Remove(dst, k)
	}
}
