package main

// The index function returns the index of the first occurrence of v in s,
// or -1 if not present.
func index[E comparable](s []E, v E) int {
	for i, vs := range s {
		if v == vs {
			return i
		}
	}
	return -1
}
