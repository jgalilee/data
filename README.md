# Data Generator

# Go Package

Install the data package with `go get github.com/jgalilee/data`. There are two
sub-packages included with this the _transactions_ package, and the _points_
package.

## Command Line Interface Options

The commandline interface leverages the _transactions_ and _points_
sub-packages to generate different sized sets of data for experimentation.
Since data uses a set seed it is possible to re-generate the data, or scale
it up / down using the same seed.

### Shared

* -type
* -size
* -measure
* -seed
* -min
* -max
* -output

### Transactions

* -catalog

### Points

* -clusters
* -dimensions
* -stddev
* -labels

## Types

### Points

The algorithm for generating the random cluster data is simple.

1. Given a value _k_ for the number of clusters, generate _k_ points in the
n-dimensional space by selecting the value of each dimension from a uniform
distribution with values in the range of min, max.
2. Once, for each cluster, generate a new point by taking that cluster as a
starting point and randomly offsetting each of dimension d in the cluster from
a normal distribution where the cluster dimension value is the mean and the
standard deviation is relative to the range min, max.

Simply put, imagine that we randomly jab a piece of paper with a red pen. After
we have done this we take a blue pen and randomly jab around each red dot in
order, trying, but not too hard, to keep the new dot in the vicinity of the red
one. 

#### Generating 1GB of Points Clustered around 5 points.

Generating 1GB of clustered point data.

```bash
./data \
	-type=points \
	-measure=GB \
	-size=1 \
	-seed=3 \
	-stddev=0.08 \
	-min=0 \
	-max=1000000 \
	-dimensions=5 \
	-clusters=10 \
	-labels=A,B,C,D,E,F,G,H,I,J
```

### Transaction

Loads in an adjacency list, this data should contain the *association* between
items. After the adjacency list has been read in the program will randomly walk
the graph generating transactions of size *[min, max)*.

*Example*

```
0\t1
0\t2
1\t0
2\t0
```

The program assumes that the id on the left of each line is the index to the
lookup array and will close the graph so that edges to nodes outside of the
indexable space are not included.

*Example*

 
```
0\t1
0\t2
1\t0
2\t0\s3
```

Node 2 will have its edge to Node 3 removed because Node 3 is not specified as
a node with edges out and is a dead end. However, nodes can be isolated or
stranded.

*Example*

 
```
0\t1
0\t2
1\t0
2\t0
3\t
```

In this case Node 3 is not referenced but also has no outgoing edges. This means
that the node is isolated, however, it is not removed from the graph and may
be the start of a random walk.

The walking algorithm is as follows ...

1. Allocate a random size for the walk *[min, max)*.
2. A nodes is selected at random from the graph.
3. This node is added to the transaction.
4. A node is selected at random by visiting one of the last selected nodes edges
at random.
  * If the node has no outgoing edges a new node is selected as in Step 1 and
  the process is repeated until an edge is visited. It is possible to visit the
  same origin node twice.
5. This node is added to the transction.
6. Once the transaction has reached the allocated size it is returned and a new
walk is started.

#### Generating 1GB of transcation data

```bash
./data \
	-type=transactions \
	-measure=GB \
	-size=1 \
	-catalog=catalog-example.txt \
	-transactions=10000 \
	-min=3 \
	-max=10 \
```

## Licence

```

The MIT License (MIT)

Copyright (c) 2014 Jack Galilee

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

```
