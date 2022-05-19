package fun_with_arrays

// checkIfExist Problem: Check If N and Its Double Exist
//
// LeetCode: https://leetcode.com/explore/learn/card/fun-with-arrays/527/searching-for-items-in-an-array/3250/
//
// Hint 1:
//	Loop from i = 0 to arr.length, maintaining in a hashTable the array elements from [0, i - 1].
// Hint 2:
//	On each step of the loop check if we have seen the element 2 * arr[i] so far or arr[i] / 2 was seen if arr[i] % 2 == 0.
func checkIfExist(arr []int) bool {
	if len(arr) == 0 {
		return false
	}

	if len(arr) == 1 {
		return false
	}

	values := map[int]int{}

	for i := 0; i < len(arr); i++ {
		values[arr[i]] = i
	}

	for i := 0; i < len(arr); i++ {
		current := arr[i]
		if current == 0 {
			if j, ok := values[current]; ok {
				if j == i {
					continue
				}
				return true
			}
		}

		if j, ok := values[current/2]; ok && current%2 == 0 {
			if j == i {
				continue
			}
			return true
		}

		if j, ok := values[2*current]; ok {
			if j == i {
				continue
			}
			return true
		}

	}

	return false
}

// validMountainArray Problem: Valid Mountain Array.
//
// LeetCode: https://leetcode.com/explore/learn/card/fun-with-arrays/527/searching-for-items-in-an-array/3251/
//
// Hint 1:
//	It's very easy to keep track of a monotonically increasing or decreasing ordering of elements.
//	You just need to be able to determine the start of the valley in the mountain and from that point onwards,
//	it should be a valley i.e. no mini-hills after that. Use this information in regards to the values in the array
//	and you will be able to come up with a straightforward solution.
func validMountainArray(arr []int) bool {
	if len(arr) == 1 {
		return false
	}

	if len(arr) == 2 {
		return false
	}

	valleyLeft := arr[0]
	valleyRight := arr[len(arr)-1]
	peak := -1
	hasPlato := false        // We don't know mountain has plato at start.
	decreasingFound := false // We don't know direction on start.

	for i := 1; i < len(arr); i++ {
		// Find top of the hill.
		if arr[i] > peak {
			peak = arr[i]
		}

		// Plato found...
		if arr[i] == arr[i-1] {
			hasPlato = true
		}

		// Going up the hill.
		if arr[i] > arr[i-1] {
			// After plato has been found
			// we should not increase.
			if hasPlato {
				return false
			}
			// If we goes down, we can't go
			// up anymore.
			if decreasingFound {
				return false
			}
		}

		// Going down the hill.
		if arr[i] < arr[i-1] {
			// decreasing
			decreasingFound = true
			if hasPlato {
				return false
			}
		}
	}

	if valleyRight >= peak {
		return false
	}

	if valleyLeft >= peak {
		return false
	}

	return true
}
