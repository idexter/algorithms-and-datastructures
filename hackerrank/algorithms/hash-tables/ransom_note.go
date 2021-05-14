package hashtables

import "fmt"

// checkMagazine solves "Ransom Note" problem.
// See: https://www.hackerrank.com/challenges/ctci-ransom-note/problem
func checkMagazine(magazine []string, note []string) {

	words := len(note)
	used := 0
	magazineMap := make(map[string]int)
	for _, v := range magazine {
		magazineMap[v]++
	}

	for _, v := range note {
		_, ok := magazineMap[v]
		if !ok {
			fmt.Println("No")
			return
		}
		magazineMap[v]--
		used++
		if ok {
			if magazineMap[v] < 0 {
				fmt.Println("No")
				return
			}
		}
	}

	if used < words {
		fmt.Println("No")
		return
	}

	fmt.Println("Yes")
}
