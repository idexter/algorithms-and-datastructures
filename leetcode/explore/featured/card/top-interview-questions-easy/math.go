package top_interview_questions_easy

import "strconv"

// fizzBuzz Problem: Fizz Buzz.
//
// LeetCode: https://leetcode.com/explore/featured/card/top-interview-questions-easy/102/math/743/
func fizzBuzz(n int) []string {
	out := make([]string, n)
	for i := 0; i < n; i++ {
		cur := i + 1
		if cur%3 == 0 && cur%5 == 0 {
			out[i] = "FizzBuzz"
			continue
		} else if cur%5 == 0 {
			out[i] = "Buzz"
			continue
		} else if cur%3 == 0 {
			out[i] = "Fizz"
			continue
		}
		out[i] = strconv.Itoa(cur)
	}
	return out
}
