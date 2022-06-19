package cookie

import (
	"fmt"
	"loyalty/internal/model"
	"net/http"
)

var session string = "Authorisation"

//По хорошему надо делать sessionStore и хранить в куках зашифрованный указатель на сессию, но не успеваю сделать так
func GetUserKeyFromCoockie(r *http.Request) (string, error) {
	token, err := r.Cookie(session)
	if err != nil {
		return "", err
	}
	if token.Value == "" {
		return "", fmt.Errorf("не найдено актуальных cookie")
	}
	userKey, err := model.Decrypt(token.Value)
	if err != nil {
		return "", err
	}
	return userKey, nil
}

func PutUserKeyToCookie(Key string) *http.Cookie {
	val, _ := model.Encrypt(Key)

	cookie := &http.Cookie{
		Name:   session,
		Value:  val,
		MaxAge: 300,
	}
	return cookie
}
