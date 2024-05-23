package gpsroutegen

func GenerateRandom(options ...GenerateRandomOption) DataList {
	var optns generateRandomOptions
	for _, opt := range options {
		opt(&optns)
	}

	if optns.start == nil {
		optns.start = ptr(RandPoint())
	}
	if optns.amount == nil {
		optns.amount = ptr(100)
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

	amounts := RandIntDistribution(*optns.amount, *optns.directionChanges)
	distances := RandIntDistribution(*optns.distance, *optns.directionChanges)

	var input []GenerateInput

	for idx, distance := range distances {
		input = append(input, GenerateInput{
			Amount: amounts[idx],
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
	amount           *int
	distance         *int
	directionChanges *int
}

// WithStart sets the starting coordinates. If not set, a random coordinate will be used.
func WithStart(start Point) GenerateRandomOption {
	return func(options *generateRandomOptions) {
		options.start = &start
	}
}

// WithAmount sets the amount of data points to generate. If not set, a value of 100 will be used.
func WithAmount(amount int) GenerateRandomOption {
	return func(options *generateRandomOptions) {
		options.amount = &amount
	}
}

// WithDistance sets the total distance of the route in meters. If not set, a random value between 1000 and 2000 will be used.
func WithDistance(distance int) GenerateRandomOption {
	return func(options *generateRandomOptions) {
		options.distance = &distance
	}
}

// WithDirectionChanges sets the amount of direction changes to generate. If not set, a random value between 0 and 5 will be used.
func WithDirectionChanges(directionChanges int) GenerateRandomOption {
	return func(options *generateRandomOptions) {
		options.directionChanges = &directionChanges
	}
}
