package main

import (
 	"github.com/jgalilee/data/points"
	"github.com/jgalilee/data/transactions"
	"os"
	"bufio"
	"fmt"
	"flag"
	"strings"
	"errors"
	"math"
)

const (
	version = "1.0"
)

type diter func() (string, bool)

// Shared flags
var command = flag.String("type", "none", "Type of data to generate.")
var size = flag.Int64("size", 10, "Number of points to generate.")
var measure = flag.String("measure", "count", "Measurement for the size.")
var seed = flag.Int64("seed", 1, "Random seed used to generate points.")
var min = flag.Int64("min", 0, "Minimum bound of a dimension coordinate.")
var max = flag.Int64("max", 100, "Maximum bound of a dimension coordinate.")
var output = flag.String("output", "/dev/stdout", "Output data destination.")

// Transcation flags
var filename = flag.String("catalog", "none", "Adjacency catalog filename.")

// Points flags
var k = flag.Int64("clusters", 2, "Number of clusters to generate.")
var d = flag.Int64("dimensions", 2, "Number of dimensions in the space.")
var stddev = flag.Float64("stddev", 0.5, "Standard deviation for generating.")
var lbls = flag.String("labels", "A,B", "Labels for the cluster centroids.")
var alternative = flag.String("alternative-output", "/dev/stdout", "Output cluster data destination")
var offset = flag.Int64("offset", 0, "Interval to write out")

// Generate a sequence of transactions given a subset of data.
func txns() diter {
	if "none" == *filename {
		err := errors.New("No catalog filename given!")
		panic(err)
	}
	file, err := os.Open(*filename)
	if nil != err {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		defer os.Exit(1)
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	catalog := transactions.LoadCatalog(scanner)
	shopper := transactions.NewShopper(*seed, int(*min), int(*max))
	iterator := func() (string, bool) {
		txn := shopper.Shop(catalog)
		return txn.String(), false
	}
	return iterator
}

// Generate a sequence of points given a set of parameters.
func pnts() diter {
	space := points.NewSpace(*seed, *d, *stddev, *min, *max)
	clusters := points.NewClusters(&space, strings.Split(*lbls, ",")...)
	i, j := int64(0), int64(0)
	iterator := func() (string, bool) {
		cluster := clusters[j]
		if i < *k {
			i += 1
			return clusters[i-1].String(), true
		}
		point := points.NewPoint(&cluster)
		j = (j+1) % *k
		return point.String(), false
	}
	return iterator
}

func byteCount(size int, measure string) int {
	var result float64
	switch measure {
	case "count":
		return size
	case "KB":
		result = 1.0
	case "MB":
		result = 2.0
	case "GB":
		result = 3.0
	default:
		result = 0.0
	}
	return int(math.Pow(1024, result)) * size
}

func check(errors ...error) {
	for _, err := range errors {
		if nil != err {
			panic(err)
		}
	}
}

func write(iterator diter) {
	// Prepare the data file.
	maxSize := byteCount(int(*size), *measure)
	offsetSize := byteCount(int(*offset), *measure)
	f1, err1 := os.Create(*output)
	f2, err2 := os.Create(*alternative)
	defer f1.Close()
	defer f2.Close()
	check(err1, err2)
	// Write out the data.
	i := 0
	var cf *os.File
	for i < maxSize + offsetSize {
		line, alternative := iterator()
		var increment int
		if (*measure == "count") {
			increment = 1
		} else {
			increment = len(line)
		}
		if !alternative {
			cf = f1
		} else {
			cf = f2
		}
		if (alternative || i > offsetSize-1) {
			_, err1 := cf.WriteString(line)
			_, err2 := cf.WriteString("\n")
			check(err1, err2)
		}
		i += increment
	}
}

func main() {
	flag.Parse()
	switch *command {
	case "points":
		write(pnts())
		defer os.Exit(1)
		break
	case "transactions":
		write(txns())
		defer os.Exit(1)
		break
	default:
		fmt.Printf("Invalid type given %v.\n", *command)
		defer os.Exit(0)
		break
	}
}
