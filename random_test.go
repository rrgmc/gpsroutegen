package gpsroutegen

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestGenerateRandom(t *testing.T) {
	data := GenerateRandom(
		WithDistance(RandRangeInt(400, 800)),
		WithStart(RandPointNear(55.953251, -3.188267, 300.0)))

	assert.Assert(t, len(data) > 0)
}
