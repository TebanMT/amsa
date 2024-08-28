package repository_impl

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (repo *ProductRepository) GetAllProductsByCompanyAndType(ctx context.Context, idCompany int, idType int) ([]entities.Product, error) {
	var products []entities.Product
	if err := repo.DB.WithContext(ctx).Where("company_id = ? && type_id = ?", idCompany, idType).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (repo *ProductRepository) RegisterProduct(ctx context.Context, product *entities.Product) error {
	return repo.DB.WithContext(ctx).Create(product).Error

}

func (repo *ProductRepository) UpdateProduct(context context.Context, product *entities.Product) error {
	return repo.DB.Model(&entities.Product{}).Where("id = ?", product.ID).Updates(product).Error
}

func (repo *ProductRepository) DeactivateProduct(context context.Context, idProduct int) error {
	result := repo.DB.Model(&entities.Product{}).Where("id = ?", idProduct).Update("activated", false).Error
	return result
}

func (repo *ProductRepository) SearchProduct(context context.Context, nameProduct string) ([]entities.Product, error) {
	var products []entities.Product
	if err := repo.DB.Where("name = ?", nameProduct).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (repo *ProductRepository) GetCountProductsByType(context context.Context, idCompany int) (map[string]int, error) {
	// Define una estructura para capturar los resultados de la consulta de agrupamiento.
	type result struct {
		Type_id string
		Count   int
	}

	var results []result
	counts := make(map[string]int)

	// Realiza la consulta: selecciona todos los productos, agrupa por tipo y cuenta las ocurrencias.
	err := repo.DB.WithContext(context).
		Model(&entities.Product{}).
		Select("type_id, COUNT(*) as count").
		Where("company_id = ? AND activated = ?", idCompany, 1).
		Group("type_id").
		Find(&results).Error

	// Verifica si hubo un error en la consulta.
	if err != nil {
		return nil, err
	}

	// Llena el mapa con los resultados obtenidos.
	for _, r := range results {
		counts[r.Type_id] = r.Count
	}

	return counts, nil
}
