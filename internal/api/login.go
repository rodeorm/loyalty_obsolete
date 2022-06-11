package api

import "net/http"

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
}
