package main

import (
	"github.com/qu1etboy/go/repository"
	"github.com/qu1etboy/go/services"
	"github.com/qu1etboy/go/transactions"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=go_db port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic("failed to connect database")
	}

	repo := repository.NewProductRepository(db)
	txManager := transactions.NewTxManager(db)
	service := services.New(repo, txManager)

	service.Run()
}
