package transactions

import (
	"sort"
	"bytes"
	"strconv"
)

const (
	idDelim = "\t"
	itemDelim = " "
)

type Transaction struct {
	Id int64
	Products []*Product
}

type ById []*Product

func (p ById) Len() int {
	return len(p)
}

func (p ById) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ById) Less(i, j int) bool {
	return p[i].Id < p[j].Id
}

func NewTransaction(id int64, products ...*Product) *Transaction {
	sort.Sort(ById(products))
	return &Transaction{Id: id, Products: products}
}

func (t Transaction) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(strconv.FormatInt(int64(t.Id), 10))
	buffer.WriteString(idDelim)
	j := len(t.Products)
	for i, p := range t.Products {
		buffer.WriteString(strconv.FormatInt(int64(p.Id), 10))
		if i + 1 < j {
			buffer.WriteString(itemDelim)
		}
	}
	return buffer.String()
}
