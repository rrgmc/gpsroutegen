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

func RandIntDistribution(amount, buckets int) []int {
	if buckets <= 0 {
		return nil
	}

	var items []int
	var total int
	for range buckets {
		item := RandRangeInt(10, 100)
		total += item
		items = append(items, item)
	}

	var totalAmount int
	var ret []int
	for idx := 0; idx < len(items)-1; idx++ {
		i := items[idx]
		pct := float64(i) / float64(total)
		curamount := int(float64(amount) * pct)
		totalAmount += curamount
		ret = append(ret, curamount)
	}
	ret = append(ret, amount-totalAmount)

	// avoid the problem that the last items usually is the bigger one.
	rand.Shuffle(len(ret), func(i, j int) { ret[i], ret[j] = ret[j], ret[i] })

	return ret
}
