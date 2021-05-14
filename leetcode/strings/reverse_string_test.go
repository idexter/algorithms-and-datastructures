package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseString(t *testing.T) {

	str := []byte("abcxyz")
	reverseString(str)
	assert.Equal(t, "zyxcba", string(str))
}
