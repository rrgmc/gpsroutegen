package gpsroutegen_test

import (
	"fmt"

	"github.com/rrgmc/gpsroutegen"
)

func ExampleGenerateRandom() {
	data := gpsroutegen.GenerateRandom(
		gpsroutegen.WithDistance(gpsroutegen.RandRangeInt(3000, 5000)),
		gpsroutegen.WithDirectionChanges(8),
		gpsroutegen.WithStart(gpsroutegen.RandPointNear(gpsroutegen.NewPoint(55.953251, -3.188267), 300.0)))

	url, err := data.ToGeoJSONIOUrl()
	if err != nil {
		panic(err)
	}
	fmt.Println("click the URL to see the route in a map")
	fmt.Println(url)
}
