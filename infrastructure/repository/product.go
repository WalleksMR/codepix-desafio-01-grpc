package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/walleksmr/codepix-desafio-01-grpc/domain/model"
)

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func (t ProductRepositoryDb) Create(product *model.Product) (*model.Product, error) {
	err := t.Db.Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (t ProductRepositoryDb) FindProducts() ([]*model.Product, error) {
	var products []*model.Product
	if err := t.Db.Find(&products).Error; err != nil {
		return nil, fmt.Errorf("Error while querying: " + err.Error())
	}
	return products, nil
}
