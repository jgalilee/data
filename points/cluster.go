package points

import (
	"fmt"
)

// Represents a cluster in the given space. It can be thought of as a labeled
// point.
type Cluster struct {
	Point
	Label string
}

// Returns a new cluster randomly generated from a uniform vector in the
// given space.
func NewCluster(space *Space, label string) Cluster {
	vector := space.UniformVector()
	point := Point{Space: space, Coords: vector}
	return Cluster{Point: point, Label: label}
}

// Returns a slice of randomly generated clusters in the given space. The
// clusters are generated from a uniform vector in the space, and the number of
// clusters returned equals the number of labels. The clusters are generated in
// the order the labels are given.
func NewClusters(space *Space, lbls ...string) []Cluster {
	result := make([]Cluster, len(lbls))
	for i := 0; i < len(result); i++ {
		result[i] = NewCluster(space, lbls[i])
	}
	return result
}

// Center label and the string representation of the centers point in the
// space.
func (c Cluster) String() string {
	return fmt.Sprintf("%v\t%v", c.Label, c.coordsAsString())
}
