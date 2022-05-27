package lab2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//	for assert equal expressions used online converter:
//	https://www.free-online-calculator-use.com/prefix-to-postfix-converter.html

func TestEasyPrefixToPostfix(t *testing.T) {
	res, err := PrefixToPostfix("+ 5 * - 4 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "5 4 2 - 3 * +", res)
	}
}

func TestMoreDigitsPrefixToPostfix(t *testing.T) {
	res, err := PrefixToPostfix("+ 52 * - 4 21 4")
	if assert.Nil(t, err) {
		assert.Equal(t, "52 4 21 - 4 * +", res)
	}
}

func TestHardPrefixToPostfix(t *testing.T) {
	res, err := PrefixToPostfix("- + - - 202 117 * 7 ^ 2 3 228 / 1337 191")
	if assert.Nil(t, err) {
		assert.Equal(t, "202 117 - 7 2 3 ^ * - 228 + 1337 191 / -", res)
	}
}

func TestExceptionPrefixToPostfix(t *testing.T) {
	_, err := PrefixToPostfix("- 2 * 2 A")
	assert.NotNil(t, err)
}

func TestEmptyPrefixToPostfix(t *testing.T) {
	_, err := PrefixToPostfix("")
	assert.NotNil(t, err)
}
