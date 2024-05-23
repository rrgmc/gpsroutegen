package gpsroutegen

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestRandDistribution(t *testing.T) {
	amount := 101
	buckets := 13

	d := RandIntDistribution(amount, buckets)

	assert.Equal(t, buckets, len(d))
	var sum int
	for _, i := range d {
		sum += i
	}
	assert.Equal(t, amount, sum)

	fmt.Println(d, sum)
}
