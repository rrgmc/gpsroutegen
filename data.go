package gpsroutegen

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

// Direction is the compass rose, naming format for readability
type Direction int

const (
	Direction_N Direction = iota
	Direction_NNE
	Direction_NE
	Direction_ENE
	Direction_E
	Direction_ESE
	Direction_SE
	Direction_SSE
	Direction_S
	Direction_SSW
	Direction_SW
	Direction_WSW
	Direction_W
	Direction_WNW
	Direction_NW
	Direction_NNW
)

const (
	compassRoseDegrees = 22.5
)

type Point struct {
	Lat, Lon float64
}

func NewPoint(lat, lon float64) Point {
	return Point{Lat: lat, Lon: lon}
}

func (p Point) ToOrbPoint() orb.Point {
	return orb.Point{p.Lon, p.Lat}
}

type Data struct {
	Point   Point
	Bearing float64
}

type DataList []Data

func (l DataList) ToGeoJSON() *geojson.FeatureCollection {
	var lines orb.LineString
	for _, g := range l {
		lines = append(lines, g.Point.ToOrbPoint())
	}

	fc := geojson.NewFeatureCollection()
	f := geojson.NewFeature(lines)
	fc.Append(f)

	return fc
}

type DataRange struct {
	Min, Max int
}

func (r DataRange) Div(amount int) DataRange {
	return DataRange{
		Min: r.Min / amount,
		Max: r.Max / amount,
	}

}

type DirectionRange struct {
	Min, Max Direction
}
