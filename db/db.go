package db

import (
	"alura-web-1/models"

	"github.com/thoas/go-funk"
)

var products = []models.Product{
	{
		Id:          1,
		Name:        "Bone",
		Description: "Bone daora",
		Price:       200.99,
	}, {
		Id:          2,
		Name:        "Camiseta",
		Description: "Camiseta loka",
		Price:       80.21,
	},
}

func FindAll() []models.Product {
	productFound := products
	return productFound
}

func FindById(id int) models.Product {
	productFound := funk.Find(products, func(p models.Product) bool {
		return p.Id == id
	}).(models.Product)

	return productFound
}

func Save(p models.Product) {
	product := models.Product{
		Id:          len(products) + 1,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}

	products = append(products, product)
}

func Update(pToUpdate models.Product) {
	index := funk.IndexOf(products, func(p models.Product) bool {
		return p.Id == pToUpdate.Id
	})

	products[index] = pToUpdate
}

func DeleteById(id int) {
	productsFiltered := funk.Filter(products, func(p models.Product) bool {
		return p.Id != id
	}).([]models.Product)

	products = productsFiltered

	// index := funk.IndexOf(products, func(p models.Product) bool {
	// 	return p.Id == id
	// })

	// products = append(products[:index], products[index+1:]...)
}
