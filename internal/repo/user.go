package repo

import (
	"context"
	"fmt"
	"loyalty/internal/model"
)

//CreateUser создает нового Пользователя
func (s *postgresStorage) CreateUser(ctx context.Context, u model.User) (dublicate bool, err error) {
	var login int
	// Проверяем, что пользователя с таким логином нет в БД
	s.DB.QueryRowContext(ctx, "SELECT login from Users WHERE Login = $1", u.Login).Scan(&login)
	if login != 0 {
		return true, nil
	}

	hashedPassword, _ := HashPassword(u.Password)

	_, err = s.DB.ExecContext(ctx, "INSERT INTO dbo.Users (Login, Password) SELECT '%s', '%s'",
		u.Login, hashedPassword)
	if err != nil {
		fmt.Println("Ошибка c запросом : ", err)
		return false, err
	}
	return false, nil
}

//AuthUser аутентифицирует Пользователя
func (s *postgresStorage) AuthUser(ctx context.Context, u model.User) (success bool, err error) {
	success, err = s.CheckPassword(ctx, u.Login, u.Password)
	if success && err == nil {
		return success, err
	}
	if err != nil {
		return false, err
	}
	return false, nil
}
