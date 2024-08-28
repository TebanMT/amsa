package use_cases

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
	"github.com/TebanMT/amsa/domain/repositories"
)

type ProductUseCases struct {
	ProductRepo     repositories.ProductRepository
	CategoryRepo    repositories.CategoryRepository
	UnitRepo        repositories.MeasurementUnitRepository
	TypeProductRepo repositories.TypeProductRepository
}

type ProductUseCaseOption func(*ProductUseCases)

func WithCategoryAndUnitRepoAndTypeRepo(catRepo repositories.CategoryRepository, unitRepo repositories.MeasurementUnitRepository, typeRepo repositories.TypeProductRepository) ProductUseCaseOption {
	return func(puc *ProductUseCases) {
		puc.CategoryRepo = catRepo
		puc.UnitRepo = unitRepo
		puc.TypeProductRepo = typeRepo
	}
}

func NewProductUseCase(productRepo repositories.ProductRepository, opts ...ProductUseCaseOption) *ProductUseCases {
	puc := &ProductUseCases{
		ProductRepo: productRepo,
	}
	for _, opt := range opts {
		opt(puc)
	}
	return puc
}

func (s *ProductUseCases) GetAllProductsByCompanyAndType(ctx context.Context, companyId int, idType int) ([]entities.Product, error) {
	return s.ProductRepo.GetAllProductsByCompanyAndType(ctx, companyId, idType)
}

func (s *ProductUseCases) RegisterProductIntoDB(
	ctx context.Context,
	product *entities.Product,
	category *entities.Category,
	unit *entities.MeasurementUnit,
	typeProduct *entities.ProductType,
) error {
	if s.CategoryRepo != nil && category.ID == 0 {
		cat, err := s.CategoryRepo.RegisterCategory(ctx, category)
		if err != nil {
			return err
		}
		product.Id_category = int(cat.ID)

	}

	if s.UnitRepo != nil && unit.ID == 0 {
		unit, err := s.UnitRepo.RegisterUnit(ctx, unit)
		if err != nil {
			return err
		}
		product.Unit_id = int(unit.ID)

	}

	if s.TypeProductRepo != nil && typeProduct.ID == 0 {
		typeProduct, err := s.TypeProductRepo.RegisterTypeProduct(ctx, typeProduct)
		if err != nil {
			return err
		}
		product.Type_id = int(typeProduct.ID)
	}

	return s.ProductRepo.RegisterProduct(ctx, product)
}

func (s *ProductUseCases) UpdateProduct(ctx context.Context, product *entities.Product) error {
	return s.ProductRepo.UpdateProduct(ctx, product)
}

func (s *ProductUseCases) DeactivateProduct(ctx context.Context, idProduct int) error {
	return s.ProductRepo.DeactivateProduct(ctx, idProduct)
}

func (s *ProductUseCases) SearchProduct(ctx context.Context, nameProduct string) ([]entities.Product, error) {
	return s.ProductRepo.SearchProduct(ctx, nameProduct)
}

func (s *ProductUseCases) GetCountProductsByType(ctx context.Context, idCompany int) (map[string]int, error) {
	return s.ProductRepo.GetCountProductsByType(ctx, idCompany)
}
