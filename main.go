package main

import (
	"log"
	"os"
	"time"

	"github.com/qu1etboy/go/repository"
	"github.com/qu1etboy/go/services"
	"github.com/qu1etboy/go/transactions"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type UserService interface {
	GetUser(id int) (any, error)
	CreateUser(user any) error
	UpdateUser(id int, user any) error
	DeleteUser(id int) error
}

type userService struct{}

type Product struct {
	ID        uint
	Title     string
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound
			Colorful:                  true,        // Disable color
		},
	)

	dsn := "host=localhost user=postgres password=postgres dbname=go_db port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}

	repo := repository.NewProductRepository(db)
	txManager := transactions.NewTxManager(db)
	service := services.New(repo, txManager)

	service.Run()
}
