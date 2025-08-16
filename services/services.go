package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/qu1etboy/go/repository"
	"github.com/qu1etboy/go/transactions"
)

type Service interface {
	Run()
}

type service struct {
	repo      repository.ProductRepository
	txManager transactions.TxManager
}

func New(repo repository.ProductRepository, txManager transactions.TxManager) Service {
	return &service{repo: repo, txManager: txManager}
}

func (s *service) Run() {
	products_a := []repository.Product{
		{Title: "Product 1", Price: 100},
		{Title: "Product 2", Price: 200},
		{Title: "Product 3", Price: 300},
	}

	products_b := []repository.Product{
		{Title: "Product 4", Price: 400},
		{Title: "Product 5", Price: 500},
		{Title: "Product 6", Price: 600},
	}

	err := s.txManager.Execute(context.Background(), func(ctx context.Context) error {
		fmt.Println("tx 1")
		err := s.repo.CreateBatch(ctx, products_a)
		if err != nil {
			return err
		}
		err = s.repo.UpdateByTitle(ctx, "Product 1", repository.Product{Title: "Product 1 Updated", Price: 200})
		if true {
			return errors.New("this will failed!!!")
		}
		return nil
	})

	if err != nil {
		fmt.Println("error executing tx", err)
	}

	err = s.txManager.Execute(context.Background(), func(ctx context.Context) error {
		fmt.Println("tx 2")
		err := s.repo.CreateBatch(ctx, products_b)
		if err != nil {
			return err
		}
		err = s.repo.UpdateByTitle(ctx, "Product 4", repository.Product{Title: "Product 4 Updated", Price: 200})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println("error executing tx", err)
	}
}
