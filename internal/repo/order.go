package repo

import (
	"context"
	"database/sql"
	"fmt"
	"loyalty/internal/model"
)

//InsertOrder создает новый заказ
func (s *postgresStorage) InsertOrderSimple(ctx context.Context, o *model.Order) (dublicateOrder, anotherUserOrder bool, err error) {
	var userLogin, orderNum string
	s.DB.QueryRowContext(ctx, "SELECT userlogin, number from Orders WHERE number = $1", o.Number).Scan(&userLogin, &orderNum)
	if userLogin != o.UserLogin && orderNum != "" {
		return true, true, nil
	}
	if orderNum != "" {
		return true, false, nil
	}
	_, err = s.DB.ExecContext(ctx, "INSERT INTO Orders (number, status, userlogin, accrual, uploadedtime) SELECT $1, $2, $3, $4, now()",
		o.Number, o.Status, o.UserLogin, o.AccrualBalls)
	if err != nil {
		fmt.Println("Ошибка c запросом : ", err)
		return false, false, err
	}
	return false, false, nil
}

func (s postgresStorage) UpdateOrder(ctx context.Context, o *model.Order) (err error) {
	_, err = s.DB.ExecContext(ctx, "UPDATE Orders"+
		"SET Status = @Status, Accrual = @Accrual, "+
		"Withdrawn = @Withdrawn "+
		"ProcessedTime = @ProcessedTime "+
		"WHERE Number = @Number",
		sql.Named("Number", o.Number),
		sql.Named("Status", o.Status),
		sql.Named("Accrual", o.AccrualBalls),
		sql.Named("Withdrawn", o.WithdrawnBalls),
		sql.Named("ProcessedTime", o.ProcessedTime),
	)
	if err != nil {
		fmt.Println("Ошибка c запросом : ", err)
		return err
	}
	return nil
}

//Получение для пользователя списка загруженных номеров заказов
func (s *postgresStorage) SelectOrders(ctx context.Context, u *model.User) (*[]model.Order, error) {

	rows, err := s.DB.QueryContext(ctx, "SELECT Number, UserLogin, Status, Accrual, UploadedTime FROM Orders WHERE UserLogin = $1", u.Login)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, err
	}

	defer rows.Close()
	orders := make([]model.Order, 0, 1)
	for rows.Next() {
		var order model.Order
		err = rows.Scan(&order.Number, &order.UserLogin, &order.Status, &order.AccrualBalls, &order.UploadedTime)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}
	return &orders, nil
}

//InsertOrder вставляет новый запрос на списание средств
func (s *postgresStorage) InsertOrder(ctx context.Context, user *model.User, order *model.Order) (dublicateOrder bool, shortage int, err error) {
	// Актуализируем данные баланса пользователя
	s.SelectUserData(ctx, user)

	// Проверяем, что заказа еще нет в базе
	var userLogin, orderNum string
	s.DB.QueryRowContext(ctx, "SELECT userlogin, number from Orders WHERE number = $1", order.Number).Scan(&userLogin, &orderNum)
	if orderNum != "" {
		return true, 0, nil
	}

	// Проверяем, что средств для списания достаточно, если нет - возвращаем информацию о дефиците баллов
	if user.Balance < order.AccrualBalls {
		shortage = user.Balance - order.AccrualBalls
		return false, shortage, err
	}

	_, err = s.DB.ExecContext(ctx, "INSERT INTO Orders (number, status, userlogin, accrual, withdrawn, uploadedtime) SELECT $1, $2, $3, $4, $5, NOW()",
		order.Number, order.Status, order.UserLogin, order.AccrualBalls, order.WithdrawnBalls)
	if err != nil {
		return false, 0, err
	}

	return false, 0, nil

}
