package factories

import (
	"math/rand"
	"time"

	"github.com/jaswdr/faker"
	"github.com/spohess/models"
)

type ProductFactory struct {
	fake faker.Faker
	rng  *rand.Rand
}

func NewProductFactory() *ProductFactory {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	return &ProductFactory{
		fake: faker.New(),
		rng:  rng,
	}
}
func (f *ProductFactory) Make() *models.Product {
	return &models.Product{
		Name:        f.fake.Lorem().Word(),
		Description: f.fake.Lorem().Paragraph(1),
		Price:       float64(f.rng.Intn(10000)) / 100,
		Quantity:    f.rng.Intn(100) + 1,
	}
}
func (f *ProductFactory) Create() *models.Product {
	product := f.Make()
	models.CreateProduct(product.Name, product.Description, product.Price, product.Quantity)
	return product
}
func (f *ProductFactory) CreateMany(count int) []*models.Product {
	products := make([]*models.Product, count)

	for i := 0; i < count; i++ {
		products[i] = f.Create()
	}

	return products
}
