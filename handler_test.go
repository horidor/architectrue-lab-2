package lab2

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEasyHandler(t *testing.T) {
	prefix := "+ 1 ^ - 2 2 3"
	postfix := "1 2 2 - 3 ^ +"
	input := strings.NewReader(prefix)
	output := bytes.NewBufferString("")

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if assert.Nil(t, err) {
		assert.Equal(t, postfix, output.String())
	}
}

func TestEmpty(t *testing.T) {
	input := strings.NewReader("")
	output := bytes.NewBufferString("")

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if assert.NotNil(t, err) {
		assert.Equal(t, "", output.String())
	}
}

func TestError(t *testing.T) {
	input := strings.NewReader("* a")
	output := bytes.NewBufferString("")

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	assert.NotNil(t, err)
}
