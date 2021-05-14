package arrays

// rotLeft rotate array left.
// See: https://www.hackerrank.com/challenges/ctci-array-left-rotation/problem
func rotLeft(a []int32, d int32) []int32 {

	for i := 0; i < int(d); i++ {
		first := a[0]
		a = a[1:]
		a = append(a, first)
	}

	return a
}
