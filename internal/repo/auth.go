package repo

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//CheckPassword проверяет пароль
func (s postgresStorage) CheckPassword(ctx context.Context, login, password string) (bool, error) {
	var passwordDB string
	err := s.DB.QueryRowContext(ctx, "SELECT Password FROM Users WHERE Login = $1; ", login).Scan(&passwordDB)
	if err != nil {
		fmt.Println("Ошибка c запросом: ", err)
		return false, nil
	}

	if CheckPasswordHash(password, passwordDB) {
		fmt.Printf("Успешная аутентификация пользователя %s", login)
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
