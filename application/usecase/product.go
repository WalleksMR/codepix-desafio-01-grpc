package usecase

import "github.com/walleksmr/codepix-desafio-01-grpc/domain/model"

type ProductUseCase struct {
	ProductRepository model.ProductRepositoryInterface
}

func (p *ProductUseCase) Create(name string, description string, price float64) (*model.Product, error) {

	productInput := &model.Product{
		Name:        name,
		Description: description,
		Price:       price,
	}

	product, err := p.ProductRepository.Create(productInput)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductUseCase) FindProducts() ([]*model.Product, error) {
	products, err := p.ProductRepository.FindProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}
