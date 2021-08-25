package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test(t *testing.T) {
	a:=1
	a++
	require.EqualValues(t, 2, a)
}
