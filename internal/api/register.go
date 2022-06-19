package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"loyalty/internal/api/cookie"
	"loyalty/internal/model"
	"net/http"
)

/*
Регистрация производится по паре логин/пароль. Каждый логин должен быть уникальным. После успешной регистрации должна происходить автоматическая аутентификация пользователя.
*/
func (h Handler) registerPost(w http.ResponseWriter, r *http.Request) {
	/*
		Регистрация производится по паре логин/пароль. Каждый логин должен быть уникальным. После успешной регистрации должна происходить автоматическая аутентификация пользователя.
		Формат запроса:
		POST /api/user/register HTTP/1.1
		Content-Type: application/json
		...

		{
		    "login": "<login>",
		    "password": "<password>"
		}
		Возможные коды ответа:
		200 — пользователь успешно зарегистрирован и аутентифицирован;
		400 — неверный формат запроса;
		409 — логин уже занят;
		500 — внутренняя ошибка сервера.
	*/
	user := model.User{}
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bodyBytes, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := context.TODO()
	isDuplicated, err := h.Storage.InsertUser(ctx, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if isDuplicated {
		w.WriteHeader(http.StatusConflict)
		return
	}
	http.SetCookie(w, cookie.PutUserKeyToCookie(user.Login))
	w.WriteHeader(http.StatusOK)
}
