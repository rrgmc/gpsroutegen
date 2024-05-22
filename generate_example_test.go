package gpsroutegen_test

import (
	"fmt"

	"github.com/rrgmc/gpsroutegen"
)

func ExampleGenerate() {
	data := gpsroutegen.Generate(gpsroutegen.NewPoint(55.953251, -3.188267),
		gpsroutegen.GenerateInput{
			Amount: 10,
			BearingRange: gpsroutegen.DirectionRange{
				Min: gpsroutegen.Direction_SSE,
				Max: gpsroutegen.Direction_SSW,
			},
			DistanceRange: gpsroutegen.DataRange{
				Min: 10 * 1000,
				Max: 15 * 1000,
			},
		},
		gpsroutegen.GenerateInput{
			Amount: 20,
			BearingRange: gpsroutegen.DirectionRange{
				Min: gpsroutegen.Direction_W,
				Max: gpsroutegen.Direction_WNW,
			},
			DistanceRange: gpsroutegen.DataRange{
				Min: 1000,
				Max: 1500,
			},
		},
	)

	enc, err := data.ToGeoJSON().MarshalJSON()
	if err != nil {
		panic(err)
	}

	// paste this output in https://geojson.io/#map=2/0/20
	fmt.Println(string(enc))
}
