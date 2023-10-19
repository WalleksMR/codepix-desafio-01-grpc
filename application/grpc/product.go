package grpc

import (
	"context"
	"strconv"

	"github.com/walleksmr/codepix-desafio-01-grpc/application/grpc/pb"
	"github.com/walleksmr/codepix-desafio-01-grpc/application/usecase"
)

type ProductGrpcService struct {
	ProductUseCase usecase.ProductUseCase
	pb.UnimplementedProductServiceServer
}

func (p *ProductGrpcService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product, err := p.ProductUseCase.Create(in.Name, in.Description, float64(in.Price))
	if err != nil {
		return nil, err
	}

	id, _ := strconv.Atoi(product.ID)

	productResponse := &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:          int32(id),
			Name:        product.Name,
			Description: product.Description,
			Price:       float32(product.Price),
		},
	}

	return productResponse, nil
}

func (p *ProductGrpcService) FindProducts(ctx context.Context, in *pb.FindProductsRequest) (*pb.FindProductsResponse, error) {
	products, err := p.ProductUseCase.FindProducts()
	if err != nil {
		return nil, err
	}

	productsResponse := &pb.FindProductsResponse{}

	for _, product := range products {
		id, _ := strconv.Atoi(product.ID)

		productResponse := &pb.Product{
			Id:          int32(id),
			Name:        product.Name,
			Description: product.Description,
			Price:       float32(product.Price),
		}

		productsResponse.Products = append(productsResponse.Products, productResponse)
	}

	return productsResponse, nil
}

func NewProductGrpcService(usecase usecase.ProductUseCase) *ProductGrpcService {
	return &ProductGrpcService{
		ProductUseCase: usecase,
	}
}
