package model

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var Products = []Product{
	{ID: "1", Name: "Laptop", Price: 1000},
	{ID: "2", Name: "Phone", Price: 500},
}
