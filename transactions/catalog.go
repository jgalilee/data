package transactions

import (
	"bufio"
	"strings"
	"strconv"
)

const (
	buffSize = 10
	delim = "\t"
)

type Catalog struct {
	Products []*Product
	numProducts int
}

func LoadCatalog(input *bufio.Scanner) *Catalog {
	currentId := -1
	var numItems int
	var similarItems []int
	result := Catalog{numProducts: -1, Products: make([]*Product, buffSize)}
	for input.Scan() {
		if err := input.Err(); nil != err {
			panic(err)
		}
		record := strings.Split(input.Text(), delim)
		var temp int64
		temp, _ = strconv.ParseInt(record[0], 10, 0)
		prodId := int(temp)
		temp, _ = strconv.ParseInt(record[1], 10, 0)
		simId := int(temp)
		if currentId != prodId {
			if -1 != currentId {
				result.add(currentId, similarItems[:numItems]...)
			}
			currentId = prodId
			similarItems = make([]int, buffSize)
			numItems = 0
		}
		similarItems[numItems] = simId
		numItems = numItems + 1
		if numItems >= len(similarItems) - 1 {
			temp := make([]int, numItems + buffSize)
			copy(temp, similarItems)
			similarItems = temp
		}
	}
	// TODO Ensure the products are sorted. Or that some state about if the
	// products need to be sorted before a lookup is done. It may be neccessary
	// to use a B-Tree index structure.
	result.closeGraph()
	return &result
}

// TODO Don't assume anything about the structure.
func (c Catalog) Lookup(id int) *Product {
	return c.Products[id]
}

func (c *Catalog) add(id int, similar ...int) *Product {
	product := NewProduct(c, id, similar...)
	c.numProducts = c.numProducts + 1
	if c.numProducts >= len(c.Products) - 1 {
		temp := make([]*Product, c.numProducts + buffSize)
		copy(temp, c.Products)
		c.Products = temp
	}
	c.Products[c.numProducts] = product
	return product
}

// Close the graph, however, don't remove isolated or stranded nodes. We are
// not concerened with these. They can count as potential spontaneous
// purchases.
func (c *Catalog) closeGraph() {
	c.Products = c.Products[:c.numProducts - 1]
	for _, product := range c.Products {
		k := 0
		ids := make([]int, len(product.SimilarIds))
		for _, id := range product.SimilarIds {
			if id < c.numProducts {
				ids[k] = id
				k += 1
			}
		}
		if k > 0 {
			k -= 1
		}
		product.SimilarIds = ids[:k]
	}
}
