package transactions

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type txKeyType string

const TxKey = txKeyType("tx")

type TxManager interface {
	Execute(ctx context.Context, fn func(ctx context.Context) error) error
}

type txManager struct {
	db *gorm.DB
}

func NewTxManager(db *gorm.DB) TxManager {
	return &txManager{db: db}
}

func (m *txManager) Execute(ctx context.Context, fn func(ctx context.Context) error) error {
	tx := ctx.Value(TxKey)
	// if tx is already set, just execute the function
	if tx != nil {
		fmt.Println("tx already set")
		return fn(ctx)
	}

	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		fmt.Println("tx not set, creating new tx")
		return fn(context.WithValue(ctx, TxKey, tx))
	})
}
