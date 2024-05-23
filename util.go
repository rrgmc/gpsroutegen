package gpsroutegen

import "math/rand/v2"

func ptr[T any](v T) *T {
	return &v
}

func RandRangeInt(min, max int) int {
	return rand.IntN(max-min) + min
}

func RandRangeFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandDirection() Direction {
	return Direction(RandRangeInt(int(Direction_N), int(Direction_NNW)))
}

func RandPoint() Point {
	return NewPoint(RandRangeFloat(-90.0, 90.0), RandRangeFloat(-180, 180))
}

func RandPointNear(lat, lon float64, maxDistance int) Point {
	return RandPointNearPoint(Point{Lat: lat, Lon: lon}, maxDistance)
}

func RandPointNearPoint(start Point, maxDistance int) Point {
	pmin := start.AddDistance(-maxDistance)
	pmax := start.AddDistance(maxDistance)

	return NewPoint(
		RandRangeFloat(pmin.Lat, pmax.Lat),
		RandRangeFloat(pmin.Lon, pmax.Lon))
}
