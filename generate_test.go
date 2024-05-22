package gpsroutegen

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestGenerate(t *testing.T) {
	data := Generate(NewPoint(55.953251, -3.188267),
		GenerateInput{
			Amount: 10,
			BearingRange: DirectionRange{
				Min: Direction_SSE,
				Max: Direction_SSW,
			},
			DistanceRange: DataRange{
				Min: 10 * 1000,
				Max: 15 * 1000,
			},
		},
		GenerateInput{
			Amount: 20,
			BearingRange: DirectionRange{
				Min: Direction_W,
				Max: Direction_WNW,
			},
			DistanceRange: DataRange{
				Min: 1000,
				Max: 1500,
			},
		},
	)

	assert.Assert(t, len(data) == 31, "length should be 31 but is %d", len(data))
}
