package repo

import (
	"context"
	"log"
	"loyalty/internal/model"
)

type Storage interface {
	InsertUser(ctx context.Context, u *model.User) (dublicate bool, err error)
	AuthUser(ctx context.Context, u *model.User) (success bool, err error)
	SelectUserData(ctx context.Context, u *model.User) error

	InsertOrderSimple(ctx context.Context, o *model.Order) (dublicateOrder, anotherUserOrder bool, err error)
	InsertOrder(ctx context.Context, user *model.User, order *model.Order) (dublicateOrder bool, shortage float64, err error)
	UpdateOrder(ctx context.Context, o *model.ExtOrder) (err error)
	SelectOrders(ctx context.Context, u *model.User) (*[]model.Order, error)
	SelectProcessingOrders() (*[]model.Order, error)

	InsertOperation(ctx context.Context, o *model.Operation) (err error)
	SelectOperations(ctx context.Context, u *model.User) (*[]model.Operation, error)
}

// NewStorage определяет место для хранения данных
func NewStorage(connectionString string) Storage {
	var storage Storage
	storage, err := InitPostgres(connectionString)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return storage
}
