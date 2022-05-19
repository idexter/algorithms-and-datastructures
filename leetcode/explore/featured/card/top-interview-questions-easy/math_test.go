package top_interview_questions_easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		in   int
		out  []string
	}{
		{"n=3", 3, []string{"1", "2", "Fizz"}},
		{"n=5", 5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{"n=15", 15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.out, fizzBuzz(tt.in))
		})
	}
}
