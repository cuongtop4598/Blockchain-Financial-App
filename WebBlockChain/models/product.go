package models

// Product contains information about a single product
type Product struct {
	ID        int    `json:"id" binding:"required"`
	Name      string `json:"name"`
	Cost      int    `json:"cost"`
	Amount    int    `json:"amount"`
	Available int    `json:"available"`
}

// We'll create a list of products
var Products = []Product{
	Product{1, "fish", 5000, 0, 10000},
	Product{2, "fish", 5000, 0, 10000},
	Product{3, "fish", 5000, 0, 10000},
	Product{4, "fish", 5000, 0, 10000},
	Product{5, "fish", 5000, 0, 10000},
	Product{6, "fish", 5000, 0, 10000},
}
