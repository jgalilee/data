package points

import (
	"fmt"
	"bytes"
	"strconv"
)

type Point struct {
	Space *Space
	Coords []float64
}

// Returns a new point pseudo-randomly generated from a cluster.
func NewPoint(cluster *Cluster) Point {
	vector := cluster.OffsetNormalVector()
	return Point{Space: cluster.Space, Coords: vector}
}

// Returns a pseudo-random vector in the clusters space which is offset from.
// the clusters position in the space.
func (p Point) OffsetNormalVector() []float64 {
	return p.Space.OffsetNormalVector(p.Coords)
}

// Returns a string representation of the point in the space.
func (p Point) String() string {
	return fmt.Sprintf("%v", p.coordsAsString())
}

// Returns a space seperated string of the coordinates of the point in the
// space.
func (p Point) coordsAsString() string {
	var buffer bytes.Buffer
	for i, j := 0, len(p.Coords); i < j; i++ {
		buffer.WriteString(strconv.FormatFloat(p.Coords[i], 'f', 4, 64))
		if i+1 < j {
			buffer.WriteString(" ")
		}
	}
	return buffer.String()
}
