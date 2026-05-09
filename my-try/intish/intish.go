package intish

// Your Intish interface goes here!

type Intish interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func IsPositive[T Intish](v T) bool {
	return v > 0
}
