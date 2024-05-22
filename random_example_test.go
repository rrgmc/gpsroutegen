package gpsroutegen_test

import (
	"fmt"

	"github.com/rrgmc/gpsroutegen"
)

func ExampleGenerateRandom() {
	data := gpsroutegen.GenerateRandom(
		gpsroutegen.WithDistance(gpsroutegen.RandRangeInt(400, 800)),
		gpsroutegen.WithStart(gpsroutegen.RandPointNear(gpsroutegen.NewPoint(55.953251, -3.188267), 300.0)))

	enc, err := data.ToGeoJSON().MarshalJSON()
	if err != nil {
		panic(err)
	}

	// paste this output in https://geojson.io/#map=2/0/20
	fmt.Println(string(enc))
}
