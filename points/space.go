package points

import (
	"math/rand"
)

// Represents an n-dimensional space in which points and clusters can reside.
type Space struct {
	rnd *rand.Rand
	Dimensions int64
	Stddev float64
	Min int64
	Max int64
}

// Returns a new n-dimensional space with a pseud-random generator seeded with
// s. The n-dimensions are bounded to values between min and max, and random
// generation can occur within the relative standard deviation of this bound
// against the point of origin.
func NewSpace(s int64, n int64, stddev float64, min int64, max int64) Space {
	src := rand.NewSource(s)
	rnd := rand.New(src)
	return Space{rnd: rnd, Dimensions: n, Stddev: stddev, Min: min, Max: max}
}

// Returns a vector in the space randomly generated from a uniform
// distribution.
func (s Space) UniformVector() []float64 {
	vector := make([]float64, s.Dimensions)
	for i, j := 0, len(vector); i < j; i++ {
		vector[i] = s.rnd.Float64() * float64((s.Max - s.Min) + s.Min)
	}
	return vector
}

// Returns a pseudo-random vector in the space which is offset from a given
// point in the space.
func (s Space) OffsetNormalVector(origin []float64) []float64 {
	vector := make([]float64, s.Dimensions)
	for i, j := 0, len(vector); i < j; i++ {
		relStddev := s.Stddev * float64((s.Max - s.Min))
		vector[i] = s.rnd.NormFloat64() * relStddev + origin[i]
	}
	return vector
}
