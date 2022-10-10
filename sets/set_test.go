package sets

import (
	"fmt"
	"testing"
)

func set[E comparable](ee ...E) map[E]struct{} {
	s := make(map[E]struct{}, len(ee))
	for _, e := range ee {
		s[e] = struct{}{}
	}
	return s
}

func to_string[E comparable](s map[E]struct{}) string {
	return fmt.Sprintf("%v", Keys(s))
}

var empty = set[int]()

func TestContains(t *testing.T) {
	tests := []struct {
		s    map[int]struct{}
		v    int
		want bool
	}{
		{empty, 0, false},
		{set(0), 0, true},
		{set(0), 1, false},
		{set(0, 1, 2), 1, true},
		{set(0, 1, 2), 3, false},
	}
	for _, tt := range tests {
		t.Run(to_string(tt.s), func(t *testing.T) {
			if got := Contains(tt.s, tt.v); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		s    map[int]struct{}
		vv   []int
		want map[int]struct{}
	}{
		{empty, []int{}, empty},
		{empty, []int{1}, set(1)},
		{set(1), []int{1}, set(1)},
		{set(1), []int{2}, set(1, 2)},
		{set(1, 2, 3), []int{2, 3, 4}, set(1, 2, 3, 4)},
	}
	for _, tt := range tests {
		t.Run(to_string(tt.s), func(t *testing.T) {
			got := Clone(tt.s)
			Insert(got, tt.vv...)
			if !Equal(tt.want, got) {
				t.Errorf("Insert() = %s, want %s", to_string(got), to_string(tt.want))
			}
		})
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		dst  map[int]struct{}
		src  map[int]struct{}
		want map[int]struct{}
	}{
		{empty, empty, empty},
		{set(1), empty, set(1)},
		{empty, set(1), set(1)},
		{set(1), set(1), set(1)},
		{set(1), set(2), set(1, 2)},
		{set(1, 2, 3), set(2, 3, 4), set(1, 2, 3, 4)},
	}
	for _, tt := range tests {
		name := to_string(tt.dst) + " ∪ " + to_string(tt.src)
		t.Run(name, func(t *testing.T) {
			got := Clone(tt.dst)
			Merge(got, tt.src)
			if !Equal(tt.want, got) {
				t.Errorf("Merge() = %s, want %s", to_string(got), to_string(tt.want))
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		dst  map[int]struct{}
		src  map[int]struct{}
		want map[int]struct{}
	}{
		{empty, empty, empty},
		{set(1), empty, set(1)},
		{empty, set(1), empty},
		{set(1), set(1), empty},
		{set(1), set(2), set(1)},
		{set(1, 2, 3), set(2, 3, 4), set(1)},
	}
	for _, tt := range tests {
		name := to_string(tt.dst) + " ∪ " + to_string(tt.src)
		t.Run(name, func(t *testing.T) {
			got := Clone(tt.dst)
			Subtract(got, tt.src)
			if !Equal(tt.want, got) {
				t.Errorf("Subtract() = %s, want %s", to_string(got), to_string(tt.want))
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		dst  map[int]struct{}
		src  map[int]struct{}
		want map[int]struct{}
	}{
		{empty, empty, empty},
		{set(1), empty, empty},
		{empty, set(1), empty},
		{set(1), set(1), set(1)},
		{set(1), set(2), empty},
		{set(1, 2, 3), set(2, 3, 4), set(2, 3)},
	}
	for _, tt := range tests {
		name := to_string(tt.dst) + " ∪ " + to_string(tt.src)
		t.Run(name, func(t *testing.T) {
			got := Clone(tt.dst)
			Intersect(got, tt.src)
			if !Equal(tt.want, got) {
				t.Errorf("Intersect() = %s, want %s", to_string(got), to_string(tt.want))
			}
		})
	}
}
