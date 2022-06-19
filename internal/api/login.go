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
Хендлер: POST /api/user/login.
Аутентификация производится по паре логин/пароль.
*/
func (h Handler) loginPost(w http.ResponseWriter, r *http.Request) {
	/*
		Формат запроса:
		POST /api/user/login HTTP/1.1
		Content-Type: application/json
		...

		{
		    "login": "<login>",
		    "password": "<password>"
		}
		Возможные коды ответа:
		200 — пользователь успешно аутентифицирован;
		400 — неверный формат запроса;
		401 — неверная пара логин/пароль;
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
	auth, err := h.Storage.AuthUser(ctx, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if auth {
		http.SetCookie(w, cookie.PutUserKeyToCookie(user.Login))
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
}
