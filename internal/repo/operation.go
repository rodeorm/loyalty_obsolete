package repo

import (
	"context"
	"fmt"
	"loyalty/internal/model"
)

func (s *postgresStorage) InsertOperation(ctx context.Context, o *model.Operation) (err error) {
	_, err = s.DB.ExecContext(ctx, "INSERT INTO Operations (OrderNumber, Sum, ProcessedTime) SELECT $1, $2, NOW()",
		o.OrderNumber, o.Sum)
	if err != nil {
		fmt.Println("Ошибка c запросом : ", err)
		return err
	}
	return nil
}
