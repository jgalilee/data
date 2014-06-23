package transactions

// Represents a.
type Product struct {
	Id int
	SimilarIds []int
	Catalog *Catalog
}

// Returns a new Product associated with the catalog.
func NewProduct(catalog *Catalog, id int, similar ...int) *Product {
	return &Product{Catalog: catalog, Id: id, SimilarIds: similar}
}
