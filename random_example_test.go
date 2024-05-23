package gpsroutegen_test

import (
	"fmt"

	"github.com/rrgmc/gpsroutegen"
)

func ExampleGenerateRandom() {
	data := gpsroutegen.GenerateRandom(
		// WithDistance sets the total distance of the route in meters. If not set, a random value between 1000 and 2000 will be used.
		gpsroutegen.WithDistance(gpsroutegen.RandRangeInt(3000, 5000)),
		// WithAmount sets the amount of data points to generate. If not set, a value of 100 will be used.
		gpsroutegen.WithAmount(gpsroutegen.RandRangeInt(90, 150)),
		// WithDirectionChanges sets the amount of direction changes to generate. If not set, a random value between 0 and 5 will be used.
		gpsroutegen.WithDirectionChanges(gpsroutegen.RandRangeInt(6, 10)),
		// WithStart sets the starting coordinates. If not set, a random coordinate will be used.
		gpsroutegen.WithStart(gpsroutegen.RandPointNear(55.953251, -3.188267, 300.0)))

	url, err := data.ToGeoJSONIOUrl()
	if err != nil {
		panic(err)
	}
	fmt.Println("click the URL to see the route in a map")
	fmt.Println(url)
}
