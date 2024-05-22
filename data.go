package gpsroutegen

import (
	"math"
	"net/url"

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
	compassRoseDegrees          = 22.5
	latitudeOneDegreeOfDistance = 111000 // metres
)

type Point struct {
	Lat, Lon float64
}

func NewPoint(lat, lon float64) Point {
	if lat < -90.0 {
		lat = -90.0
	}
	if lat > 90.0 {
		lat = 90.0
	}
	if lon < -180.0 {
		lon = -180.0
	}
	if lon > 180.0 {
		lon = 180.0
	}

	return Point{Lat: lat, Lon: lon}
}

func (p Point) ToOrbPoint() orb.Point {
	return orb.Point{p.Lon, p.Lat}
}

func (p Point) AddDistance(meters int) Point {
	coef := float64(meters) / float64(latitudeOneDegreeOfDistance)

	new_lat := p.Lat + coef

	// pi / 180 ~= 0.01745
	new_long := p.Lon + coef/math.Cos(p.Lat*math.Pi/180.0)

	return NewPoint(new_lat, new_long)
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

func (l DataList) ToGeoJSONIOUrl() (string, error) {
	enc, err := l.ToGeoJSON().MarshalJSON()
	if err != nil {
		return "", err
	}

	u, err := url.Parse("https://geojson.io")
	if err != nil {
		return "", err
	}

	u.Fragment = "data=data:application/json," + string(enc)
	return u.String(), nil
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
