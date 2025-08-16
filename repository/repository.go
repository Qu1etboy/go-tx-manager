package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint
	Title     string
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductRepository struct {
	baseRepository
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{baseRepository{db: db}}
}

func (r *ProductRepository) Create(ctx context.Context, product Product) error {
	return r.getDB(ctx).Create(&product).Error
}

func (r *ProductRepository) CreateBatch(ctx context.Context, products []Product) error {
	return r.getDB(ctx).CreateInBatches(&products, 100).Error
}

func (r *ProductRepository) UpdateByTitle(ctx context.Context, title string, product Product) error {
	return r.getDB(ctx).Model(&Product{}).Where("title = ?", title).Updates(product).Error
}

func (r *ProductRepository) Get(ctx context.Context, id uint) (Product, error) {
	var product Product
	if err := r.getDB(ctx).First(&product, id).Error; err != nil {
		return Product{}, err
	}
	return product, nil
}
