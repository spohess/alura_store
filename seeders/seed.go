package main

import (
	"github.com/spohess/factories"
)

func main() {
	productFactory := factories.NewProductFactory()
	productFactory.CreateMany(50)
}
