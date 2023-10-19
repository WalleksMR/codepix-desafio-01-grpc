package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type ProductRepositoryInterface interface {
	Create(product *Product) (*Product, error)
	FindProducts() ([]*Product, error)
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Product struct {
	ID          string  `json:"id" gorm:"type:uuid;primary_key" valid:"uuid"`
	Name        string  `json:"name" gorm:"type:varchar(255)" valid:"notnull"`
	Description string  `json:"description" gorm:"type:varchar(255)" valid:"notnull"`
	Price       float64 `json:"price" gorm:"type:varchar(255)" valid:"notnull"`
}

func (product *Product) isValid() error {
	_, err := govalidator.ValidateStruct(product)
	if err != nil {
		return err
	}
	return nil
}

func NewProduct(description string, name string, price float64) (*Product, error) {
	product := Product{
		Name:        name,
		Description: description,
		Price:       price,
	}
	product.ID = uuid.NewV4().String()
	err := product.isValid()
	if err != nil {
		return nil, err
	}
	return &product, nil
}
