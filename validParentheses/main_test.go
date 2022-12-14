package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValid(t *testing.T) {
	c := require.New(t)

	arrTest := []struct {
		input  string
		output bool
	}{
		{
			input:  "()",
			output: true,
		},
		{
			input:  "()[]{}",
			output: true,
		},
		{
			input:  "(]",
			output: false,
		},
		{
			input:  "{[]}",
			output: true,
		},
		{
			input:  "([{])",
			output: false,
		},
		{
			input:  "{{[{}]}}",
			output: true,
		},
		{
			input:  "{[{]}",
			output: false,
		},
	}

	for _, e := range arrTest {
		c.Equal(e.output, isValid(e.input), fmt.Sprintf("Input: %s", e.input))
	}
}
