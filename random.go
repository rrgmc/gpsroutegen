package gpsroutegen

func GenerateRandom(options ...GenerateRandomOption) DataList {
	var optns generateRandomOptions
	for _, opt := range options {
		opt(&optns)
	}

	if optns.start == nil {
		optns.start = ptr(RandPoint())
	}
	if optns.distance == nil {
		optns.distance = ptr(RandRangeInt(1000, 2000))
	}
	if optns.directionChanges == nil {
		optns.directionChanges = ptr(RandRangeInt(0, 5))
	}

	if *optns.distance < 1 {
		*optns.distance = 100
	}
	if *optns.directionChanges < 1 {
		*optns.directionChanges = 1
	}

	distanceLeft := *optns.distance
	var distances []int
	for range *optns.directionChanges - 1 {
		newDistance := distanceLeft / (*optns.directionChanges - len(distances))
		distanceLeft -= newDistance
		distances = append(distances, newDistance)
	}
	distances = append(distances, distanceLeft)

	var input []GenerateInput

	for _, distance := range distances {
		input = append(input, GenerateInput{
			Amount: 100,
			BearingRange: DirectionRange{
				Min: RandDirection(),
				Max: RandDirection(),
			},
			DistanceRange: DataRange{
				Min: distance,
				Max: distance,
			},
		})
	}

	return Generate(*optns.start, input...)
}

type GenerateRandomOption func(options *generateRandomOptions)

type generateRandomOptions struct {
	start            *Point
	distance         *int
	directionChanges *int
}

func WithStart(start Point) GenerateRandomOption {
	return func(options *generateRandomOptions) {
		options.start = &start
	}
}

func WithDistance(distance int) GenerateRandomOption {
	return func(options *generateRandomOptions) {
		options.distance = &distance
	}
}

func WithDirectionChanges(directionChanges int) GenerateRandomOption {
	return func(options *generateRandomOptions) {
		options.directionChanges = &directionChanges
	}
}
