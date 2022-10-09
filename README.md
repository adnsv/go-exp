# go-exp

A little GO library that provides additional functionality for generic maps and sets

Since the introduction of generics in GO, its experimental repository has a
couple of packages that helps to write nicer code for slices and maps:

- https://pkg.go.dev/golang.org/x/exp/slices
- https://pkg.go.dev/golang.org/x/exp/maps

I felt like adding a bit more to it, therefore this add-on library was created.

Here are the features:

- `github.com/adnsv/go-exp/maps` package 
  - flattening maps into slices of key-value pairs
  - sorting key-value pairs by key and by value
  - one-liner `range for` loops for key or value ordered iterating over existing maps

- `github.com/adnsv/go-exp/sets` package
  - implements intersect, union, difference, and other set-like functionality 
    for `map[K comparable]struct{}`