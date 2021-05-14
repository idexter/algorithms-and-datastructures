package constructive

import (
	"fmt"
	"math"
)

// minimumBribes solves New year Chaos problem.
// See: https://www.hackerrank.com/challenges/new-year-chaos/problem
func minimumBribes(q []int32) {
	bribes := 0
	for i := 0; i < len(q); i++ {
		current := q[i]
		inOrder := i + 1
		if int(current) != inOrder {
			if int(current)-inOrder > 2 {
				fmt.Printf("Too chaotic\n")
				return
			}

		}
		max := math.Max(0, float64(q[i]-3))
		for j := int(max); j <= i; j++ {
			if current < q[j] {
				bribes++
			}
		}
	}

	fmt.Printf("%d\n", bribes)
}
