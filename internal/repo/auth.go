package repo

import (
	"context"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//CheckPassword проверяет пароль
func (s postgresStorage) CheckPassword(ctx context.Context, login int, password string) (bool, error) {
	var passwordDB string
	err := s.DB.QueryRowContext(ctx, "SELECT Password FROM Users WHERE Login = @login; ", sql.Named("login", login)).Scan(passwordDB)
	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return false, err
	}

	if CheckPasswordHash(password, passwordDB) {
		fmt.Printf("Успешная аутентификация пользователя %d", login)
		return true, nil
	}
	return false, nil
}

//HashPassword хэширует пароль
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash проверят соответствие пароля и хэша
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
