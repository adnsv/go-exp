package maps

import (
	"fmt"
	"strings"

	"github.com/adnsv/go-exp/sets"
)

func ExampleInsert() {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		6: "six",
	}

	type pair = Pair[int, string]
	Insert(m, &pair{3, "THREE"}, &pair{4, "FOUR"})

	m2 := map[int]string{
		1: "UNO",
		2: "DOS",
		3: "TRES",
		5: "CINCO",
	}
	Insert(m, Pairs(m2)...)
	for _, p := range SortedByKey(m) {
		fmt.Printf("%d: %s\n", p.Key, p.Val)
	}
	// Output:
	// 1: one
	// 2: two
	// 3: three
	// 4: FOUR
	// 5: CINCO
	// 6: six
}

func ExampleInsertOrOverwrite() {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		6: "six",
	}

	type pair = Pair[int, string]
	InsertOrOverwrite(m, &pair{3, "THREE"}, &pair{4, "FOUR"})

	m2 := map[int]string{
		1: "UNO",
		2: "DOS",
		3: "TRES",
		5: "CINCO",
	}
	InsertOrOverwrite(m, Pairs(m2)...)
	for _, p := range SortedByKey(m) {
		fmt.Printf("%d: %s\n", p.Key, p.Val)
	}
	// Output:
	// 1: UNO
	// 2: DOS
	// 3: TRES
	// 4: FOUR
	// 5: CINCO
	// 6: six
}

func ExampleSortedByKey() {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
	}

	for _, p := range SortedByKey(m) {
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

	for _, p := range SortedByKey(Sliced(m, s)) {
		fmt.Printf("%d: %s\n", p.Key, p.Val)
	}
	// Output:
	// 2: two
	// 4: four
}

func ExampleMerge() {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
	}
	m2 := map[int]string{
		2: "TWO",
		3: "THREE",
		4: "four",
		5: "five",
	}

	conflicts := Merge(m, m2)

	fmt.Printf("\nMERGED\n")
	for _, p := range SortedByKey(m) {
		fmt.Printf("%d: %s\n", p.Key, p.Val)
	}
	fmt.Printf("\nCONFLICTS\n")
	for _, p := range SortedByKey(conflicts) {
		fmt.Printf("%d: %s\n", p.Key, p.Val)
	}
	// Output:
	//
	// MERGED
	// 1: one
	// 2: two
	// 3: three
	// 4: four
	// 5: five
	//
	// CONFLICTS
	// 2: TWO
	// 3: THREE
}

func ExampleInverted() {
	m := map[string]int{
		"one":                      1,
		"two":                      2,
		"three":                    3,
		"fourty two":               42,
		"the answer to everything": 42,
	}

	inverted, duplicates := Inverted(m)

	fmt.Printf("\nINVERTED\n")
	for _, p := range SortedByKey(inverted) {
		fmt.Printf("%d: %s\n", p.Key, p.Val)
	}
	fmt.Printf("\nDUPLICATES\n")
	for v := range duplicates {
		matching_keys := MatchValue(m, v)
		as_slice := sets.Sorted(matching_keys)
		fmt.Printf("%d: %s\n", v, strings.Join(as_slice, ", "))
	}
	// Output:
	//
	// INVERTED
	// 1: one
	// 2: two
	// 3: three
	//
	// DUPLICATES
	// 42: fourty two, the answer to everything
}
