package warm_up

// simpleArraySum Problem: Simple Array Sum.
//
// HackerRank: https://www.hackerrank.com/challenges/simple-array-sum/problem
func simpleArraySum(ar []int32) int32 {
	var sum int32
	for _, v := range ar {
		sum += v
	}
	return sum
}
