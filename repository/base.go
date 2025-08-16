package repository

import (
	"context"
	"fmt"

	"github.com/qu1etboy/go/transactions"
	"gorm.io/gorm"
)

type baseRepository struct {
	db *gorm.DB
}

// getDB returns the database connection for the given context.
// If the context has a transaction, it returns the transaction database connection.
// Otherwise, it returns the main database connection.
func (r *baseRepository) getDB(ctx context.Context) *gorm.DB {
	tx := ctx.Value(transactions.TxKey)
	if tx != nil {
		fmt.Println("tx found")
		return tx.(*gorm.DB).WithContext(ctx)
	}
	fmt.Println("tx not found")
	return r.db.WithContext(ctx)
}
