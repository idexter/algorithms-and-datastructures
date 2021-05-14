package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseString(t *testing.T) {

	str1 := []byte("hello")
	reverseString(str1)
	assert.Equal(t, "olleh", string(str1))

	str2 := []byte("Hannah")
	reverseString(str2)
	assert.Equal(t, "hannaH", string(str2))
}
