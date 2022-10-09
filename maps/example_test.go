package maps

import (
	"fmt"
)

func ExampleStableSortedByKey() {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
	}

	for _, p := range StableSortedByKey(m) {
		fmt.Printf("%d: %s\n", p.Key, p.Val)
	}
	// Output:
	// 1: one
	// 2: two
	// 3: three
	// 4: four
}

func ExampleStableSortedByVal() {
	m := map[int]string{
		1: "D",
		2: "C",
		3: "B",
		4: "A",
	}

	for _, p := range StableSortedByVal(m) {
		fmt.Printf("%d: %s\n", p.Key, p.Val)
	}
	// Output:
	// 4: A
	// 3: B
	// 2: C
	// 1: D
}

func ExampleSliced() {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
	}

	s := map[int]struct{}{
		2: {},
		4: {},
	}

	for _, p := range StableSortedByKey(Sliced(m, s)) {
		fmt.Printf("%d: %s\n", p.Key, p.Val)
	}
	// Output:
	// 2: two
	// 4: four
}
