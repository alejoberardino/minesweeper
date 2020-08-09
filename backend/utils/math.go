package utils

// Weird that I have to actually write this and it's not native...
func Min(n1 int, n2 int) int {
	if n1 > n2 {
		return n2
	}
	return n1
}
func Max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
