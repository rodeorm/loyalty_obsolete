package repo

import (
	"context"
	"database/sql"
	"fmt"
	"loyalty/internal/model"
)

//CreateOrder создает новый заказ
func (s *postgresStorage) CreateOrder(ctx context.Context, o model.Order) (dublicateOrder, anotherUserOrder bool, err error) {
	var userLogin, orderNum int
	// Проверяем, что заказа еще не создавалось
	s.DB.QueryRowContext(ctx, "SELECT userlogin, ID from Users WHERE ID = $1", orderNum).Scan(&userLogin, &orderNum)
	if userLogin != o.User.Login && orderNum != 0 {
		return true, true, nil
	}
	if orderNum != 0 {
		return true, false, nil
	}

	_, err = s.DB.ExecContext(ctx, "INSERT INTO dbo.Orders (ID, OrderStatusID, UserLogin, Accrual) SELECT @ID, @OrderStatusID, @UserLogin, @Accrual",
		sql.Named("ID", o.Number),
		sql.Named("OrderStatusID", o.Status.ID),
		sql.Named("UserLogin", o.User.Login),
		sql.Named("Accrual", o.Accrual),
	)
	if err != nil {
		fmt.Println("Ошибка c запросом : ", err)
		return false, false, err
	}
	return false, false, nil
}
