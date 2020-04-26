package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_formatAddress(t *testing.T) {
	assert.EqualValues(t,
		"000000000000000000000000dda250dd2646e02ee403da26eb7065dedafb79fd",
		formatAddress("dda250dd2646e02ee403da26eb7065dedafb79fd"))
}
