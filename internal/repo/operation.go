package repo

import (
	"context"
	"fmt"
	"sort"

	"github.com/rodeorm/loyalty/internal/model"
)

func (s *postgresStorage) InsertOperation(ctx context.Context, o *model.Operation) (err error) {
	_, err = s.DB.ExecContext(ctx, "INSERT INTO Operations (OrderNumber, Sum, UserLogin, ProcessedTime) SELECT $1, $2, $3, NOW()",
		o.OrderNumber, o.Sum, o.UserLogin)
	if err != nil {
		fmt.Println("Ошибка c запросом : ", err)
		return err
	}

	return nil
}

// Получение для пользователя списка загруженных номеров заказов
func (s *postgresStorage) SelectOperations(ctx context.Context, u *model.User) (*[]model.Operation, error) {

	rows, err := s.DB.QueryContext(ctx, "SELECT OrderNumber, Sum, ProcessedTime FROM Operations WHERE UserLogin = $1", u.Login)
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, err
	}
	defer rows.Close()
	operations := make([]model.Operation, 0, 1)
	for rows.Next() {
		var operation model.Operation
		err = rows.Scan(&operation.OrderNumber, &operation.Sum, &operation.ProcessedTime)
		if err != nil {
			return nil, err
		}

		operations = append(operations, operation)
	}
	sort.Slice(operations, func(i, j int) bool {
		return operations[i].ProcessedTime.Before(operations[j].ProcessedTime)
	})
	return &operations, nil
}
