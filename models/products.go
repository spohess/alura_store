package models

import (
	"time"

	"github.com/spohess/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func CreateProduct(name string, description string, price float64, quantity int) {
	db := db.DbConnect()

	query, err := db.Prepare("INSERT INTO products (name, description, price, quantity, created_at, updated_at) VALUES ($1, $2, $3, $4, now(), now())")
	if err != nil {
		panic("Erro de criação ed insert de produto: " + err.Error())
	}

	query.Exec(name, description, price, quantity)

	defer db.Close()
}

func DeleteProduct(id int) {
	db := db.DbConnect()

	query, err := db.Prepare("DELETE FROM products WHERE id=$1")
	if err != nil {
		panic("Erro na remoção do produto: " + err.Error())
	}

	query.Exec(id)

	defer db.Close()
}

func GetAllProducts() []Product {
	db := db.DbConnect()
	query, err := db.Query("SELECT * FROM products ORDER BY id DESC")
	if err != nil {
		panic("Erro ao coletar produtos: " + err.Error())
	}
	p := Product{}
	products := []Product{}

	for query.Next() {
		var id, quantity int
		var name, description string
		var price float64
		var created_at, updated_at time.Time

		err = query.Scan(&id, &name, &description, &price, &quantity, &created_at, &updated_at)
		if err != nil {
			panic("Erro ao atribuir query: " + err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity
		p.CreatedAt = created_at
		p.UpdatedAt = updated_at

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func GetProductById(id int) Product {
	db := db.DbConnect()
	query, err := db.Query("SELECT * FROM products WHERE id=$1", id)
	if err != nil {
		panic("Erro ao coletar o produto: " + err.Error())
	}
	product := Product{}

	for query.Next() {
		var id, quantity int
		var name, description string
		var price float64
		var created_at, updated_at time.Time

		err = query.Scan(&id, &name, &description, &price, &quantity, &created_at, &updated_at)
		if err != nil {
			panic("Erro ao atribuir query: " + err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity
		product.CreatedAt = created_at
		product.UpdatedAt = updated_at
	}
	defer db.Close()
	return product
}
