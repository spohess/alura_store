package models

import "github.com/spohess/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {
	db := db.DbConnect()
	query, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic("Erro ao coletar produtos: " + err.Error())
	}
	p := Product{}
	products := []Product{}

	for query.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = query.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic("Erro ao atribuir query: " + err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}
