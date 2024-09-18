package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rodeorm/loyalty/internal/api/cookie"
	"github.com/rodeorm/loyalty/internal/model"
)

/*
Получение информации о выводе средств
Хендлер: GET /api/user/balance/withdrawals.
Хендлер доступен только авторизованному пользователю. Факты выводов в выдаче должны быть отсортированы по времени вывода от самых старых к самым новым. Формат даты — RFC3339.
*/
func (h Handler) balanceWithdrawalsGet(w http.ResponseWriter, r *http.Request) {
	/*
		Получение информации о выводе средств
		Хендлер: GET /api/user/balance/withdrawals.
		Хендлер доступен только авторизованному пользователю. Факты выводов в выдаче должны быть отсортированы по времени вывода от самых старых к самым новым. Формат даты — RFC3339.
		Формат запроса:
		GET /api/user/withdrawals HTTP/1.1
		Content-Length: 0
		Возможные коды ответа:
		200 — успешная обработка запроса.
		Формат ответа:
		  200 OK HTTP/1.1
		  Content-Type: application/json
		  ...

		  [
		      {
		          "order": "2377225624",
		          "sum": 500,
		          "processed_at": "2020-12-09T16:09:57+03:00"
		      }
		  ]

		204 — нет ни одного списания.
		401 — пользователь не авторизован.
		500 — внутренняя ошибка сервера.
	*/
	ctx := context.TODO()
	userKey, err := cookie.GetUserKeyFromCoockie(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user := model.User{Login: userKey}
	fmt.Println(user)
	operationHistory, err := h.Storage.SelectOperations(ctx, &user)

	if err != nil {
		fmt.Println("Проблемы с получением истории списаний", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	bodyBytes, err := json.Marshal(operationHistory)
	if err != nil {
		fmt.Println("Проблемы при маршалинге истории списаний", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(bodyBytes))

}
