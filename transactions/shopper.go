package transactions

import (
	"math/rand"
)

// Represents a shopper who can conduct transactions of a minimum and maximum
// size.
type Shopper struct {
	Min int
	Max int
	rnd *rand.Rand
	conducted int64
}

// Returns a new shopper which will randomly shop from the catalog with a cart
// size no smaller than min, and no larger than max.
func NewShopper(s int64, min int, max int) Shopper {
	src := rand.NewSource(s)
	rnd := rand.New(src)
	return Shopper{Min: min, Max: max, rnd: rnd, conducted: int64(-1)}
}

// Conducts a new transction from the shoppers maximum and minimum cart size.
// The first item is initially randomly selected from the catalog using a
// uniform distribution. Subsequent items are randomly selected from the inital
// items similar items list.
func (s *Shopper) Shop(c *Catalog) *Transaction {
	txnLen := s.rnd.Intn(s.Max - s.Min) + s.Min;
	txnCart := make([]*Product, txnLen)
	browsing := c.Lookup(s.rnd.Intn(len(c.Products) - 1))
	for i := 0; i < txnLen; i++ {
		txnCart[i] = browsing
		var nextId int
		// If a dead-end has been reached re-adjust.
		// TODO Don't readjust to a visited node.
		for 0 == len(browsing.SimilarIds) {
			browsing = c.Lookup(s.rnd.Intn(len(c.Products) - 1))
		}
		// Is there only one item to choose from? If so choose it and re-adjust
		// if necessary next time.
		idx := 0
		if 1 < len(browsing.SimilarIds) {
			idx = s.rnd.Intn(len(browsing.SimilarIds) - 1)
		}
		nextId = browsing.SimilarIds[idx]
		browsing = c.Lookup(nextId)
	}
	s.conducted += 1
	return NewTransaction(s.conducted, txnCart...)
}
