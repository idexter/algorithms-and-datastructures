package arrays

// hourglassSum Problem: 2D Array - DS.
//
// HackerRank: https://www.hackerrank.com/challenges/2d-array/problem
func hourglassSum(arr [][]int32) int32 {

	calcHourGlass := func(arr [][]int32, x, y int) int32 {
		return arr[0+x][0+y] + arr[0+x][1+y] + arr[0+x][2+y] + arr[1+x][1+y] + arr[2+x][0+y] + arr[2+x][1+y] + arr[2+x][2+y]
	}

	var res int32 = -63
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			hg := calcHourGlass(arr, x, y)
			if hg >= res {
				res = hg
			}
		}
	}

	return res
}
