package fun_with_arrays

// duplicateZeros Problem: Duplicate Zeros.
//
// LeetCode: https://leetcode.com/explore/learn/card/fun-with-arrays/525/inserting-items-into-an-array/3245/
//
// Hint 1:
//	This is a great introductory problem for understanding and working with the concept of in-place operations.
//	The problem statement clearly states that we are to modify the array in-place.
//	That does not mean we cannot use another array. We just don't have to return anything.
// Hint 2:
//	A better way to solve this would be without using additional space.
//	The only reason the problem statement allows you to make modifications in place is that it hints at avoiding
//	any additional memory.
// Hint 3:
//	The main problem with not using additional memory is that we might override elements due to the zero duplication
//	requirement of the problem statement. How do we get around that?
// Hint 4:
//	If we had enough space available, we would be able to accommodate all the elements properly.
//	The new length would be the original length of the array plus the number of zeros.
//	Can we use this information somehow to solve the problem?
func duplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			for k := len(arr) - 1; k >= i; k-- {
				if k == 0 {
					continue
				}
				arr[k] = arr[k-1]
			}
			arr[i] = 0
			i++
		}
	}
}

// merge Problem: Merge Sorted Array.
//
// LeetCode: https://leetcode.com/explore/learn/card/fun-with-arrays/525/inserting-items-into-an-array/3253/
//
// Hint 1:
//	If we had enough space available, we would be able to accommodate all the elements properly.
//	The new length would be the original length of the array plus the number of zeros.
//	Can we use this information somehow to solve the problem?
// Hint 2:
//	If you simply consider one element each at a time from the two arrays and make a decision and proceed accordingly,
//	you will arrive at the optimal solution.
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums3 := make([]int, m+n)
	for i := 0; i < m; i++ {
		nums3[i] = nums1[i]
	}

	i, j, k := 0, 0, 0
	for i < m && j < n {
		if nums3[i] < nums2[j] {
			nums1[k] = nums3[i]
			k++
			i++
		} else {
			nums1[k] = nums2[j]
			k++
			j++
		}
	}

	for i < m {
		nums1[k] = nums3[i]
		k++
		i++
	}

	for j < n {
		nums1[k] = nums2[j]
		k++
		j++
	}
}
