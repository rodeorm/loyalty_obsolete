package repo

import (
	"context"
	"fmt"

	"github.com/rodeorm/loyalty/internal/model"
)

// InsertUser создает нового Пользователя
func (s *postgresStorage) InsertUser(ctx context.Context, u *model.User) (dublicate bool, err error) {
	var login string
	// Проверяем, что пользователя с таким логином нет в БД
	s.DB.QueryRowContext(ctx, "SELECT login from Users WHERE Login = $1", u.Login).Scan(&login)
	if login != "" {
		return true, nil
	}

	hashedPassword, _ := HashPassword(u.Password)

	_, err = s.DB.ExecContext(ctx, "INSERT INTO Users (Login, Password) SELECT $1, $2", u.Login, hashedPassword)
	if err != nil {
		fmt.Println("Ошибка c запросом : ", err)
		return false, err
	}
	return false, nil
}

// AuthUser аутентифицирует Пользователя
func (s *postgresStorage) AuthUser(ctx context.Context, u *model.User) (success bool, err error) {
	success, err = s.CheckPassword(ctx, u.Login, u.Password)
	if success && err == nil {
		return success, err
	}
	if err != nil {
		return false, err
	}
	return false, nil
}

// SelectUserData актуализирует для Пользователя данные о текущей сумме баллов лояльности, а также сумме использованных за весь период регистрации баллов
func (s *postgresStorage) SelectUserData(ctx context.Context, u *model.User) error {
	queryAccrual := "SELECT COALESCE(SUM(Accrual),0)" +
		" FROM Users AS u " +
		" LEFT JOIN Orders AS o ON o.userlogin = u.login AND o.Status = 'PROCESSED' " +
		" WHERE u.login = $1"
	queryWithdraw := "SELECT COALESCE(SUM(op.Sum),0) AS Withdrawn" +
		" FROM Users AS u " +
		" LEFT JOIN Operations AS op ON op.userLogin = u.login " +
		" WHERE u.login = $1"
	err := s.DB.QueryRowContext(ctx,
		queryAccrual, u.Login).Scan(&u.Balance)
	if err != nil {
		return err
	}
	err = s.DB.QueryRowContext(ctx,
		queryWithdraw, u.Login).Scan(&u.Withdrawn)
	if err != nil {
		return err
	}
	u.Balance = u.Balance - u.Withdrawn
	fmt.Println("Обновили данные по пользователю: ", u.Login, "Баланс: ", u.Balance, ". Списания: ", u.Withdrawn)
	return nil
}
