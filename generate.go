package gpsroutegen

import (
	"math"
	"math/rand"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geo"
)

type GenerateInput struct {
	Amount        int
	BearingRange  DirectionRange
	DistanceRange DataRange
}

func Generate(start Point, input ...GenerateInput) DataList {
	line := DataList{{
		Point:   start,
		Bearing: 0,
	}}

	for _, i := range input {
		for range i.Amount {
			newPoint := generateNewLocation(line[len(line)-1].Point,
				i.BearingRange,
				i.DistanceRange.Div(i.Amount))
			line = append(line, newPoint)
		}
	}

	return line
}

// generateNewLocation returns a new point in the range of direction and
// distance. It is meant to build non-repetitive but predictable GPS tracks, to
// help generate test input cases.
//
// It's also meant to be readable code.
func generateNewLocation(start Point, direction DirectionRange, distance DataRange) Data {

	var (
		latitudeOneDegreeOfDistance = 111000 // metres
		newPoint                    Point    // []float64{Long, Lat}

		// convert from degrees to radians
		deg2rad = func(d float64) float64 { return d * math.Pi / 180 }
	)

	// Use trigonometry of a right-angled triangle to solve the distances on the ground.
	// The hypotenuse is our desired distance to travel,  and one angle
	// is our desired bearing.
	//
	// now work out the vertical (longitude) and horizontal (latitude) sides in
	// distance units.
	hyp := (rand.Float64() * float64(distance.Max-distance.Min)) + float64(distance.Min)

	// Get the compass bearing in degrees, with a little randomness between the
	// general direction. Non-linear tracks are easier to troubleshoot visually.
	bearingMin := float64(direction.Min) * compassRoseDegrees
	bearingMax := float64(direction.Max) * compassRoseDegrees
	angle := (rand.Float64() * (bearingMax - bearingMin)) + bearingMin

	// Calulate the other side lengths using SOH CAH TOA. The Go math package
	// works in radians
	adj := math.Cos(deg2rad(angle)) * hyp // adjacent side of angle
	opp := math.Sin(deg2rad(angle)) * hyp // opposite side of angle

	// Each degree change in every latitude equates to ~111 km on the ground. So
	// now find the degree change required for the length of adj
	latitudeDelta := (1.0 / float64(latitudeOneDegreeOfDistance)) * adj
	newPoint.Lat = start.Lat + latitudeDelta

	// Distance on the ground for each degree of longitude changes depending on
	// latitude because the earth is not perfectly spherical. So we need to
	// calculate the distance of one degree longitude for our current latitude.
	p1 := orb.Point{1.0, start.Lat}
	p2 := orb.Point{2.0, start.Lat}
	longitudeOneDegreeOfDistance := geo.Distance(p1, p2) // returns metres

	// Now we can use this value to calculate the longitude degree change
	// required to move opp distance (in a horizontal straight line) at this
	// latitude.
	longitudeDelta := (1.0 / longitudeOneDegreeOfDistance) * opp

	// The new point is a vertical and horizontal shift to arrive at hyp
	// distance from the start point on the required bearing.
	newPoint.Lon = start.Lon + longitudeDelta

	return Data{Point: newPoint, Bearing: angle}
}
