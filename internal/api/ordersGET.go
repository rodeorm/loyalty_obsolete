package api

import (
	"context"
	"encoding/json"
	"fmt"
	"loyalty/internal/api/cookie"
	"loyalty/internal/model"
	"net/http"
)

/*
Получение списка загруженных номеров заказов
Хендлер: GET /api/user/orders.
Хендлер доступен только авторизованному пользователю. Номера заказа в выдаче должны быть отсортированы по времени загрузки от самых старых к самым новым. Формат даты — RFC3339.
*/
func (h Handler) ordersGet(w http.ResponseWriter, r *http.Request) {

	/*
	   Доступные статусы обработки расчётов:
	   NEW — заказ загружен в систему, но не попал в обработку;
	   PROCESSING — вознаграждение за заказ рассчитывается;
	   INVALID — система расчёта вознаграждений отказала в расчёте;
	   PROCESSED — данные по заказу проверены и информация о расчёте успешно получена.
	   Формат запроса:
	   GET /api/user/orders HTTP/1.1
	   Content-Length: 0
	   Возможные коды ответа:
	   200 — успешная обработка запроса.
	   Формат ответа:
	     200 OK HTTP/1.1
	     Content-Type: application/json
	     ...

	     [
	         {
	             "number": "9278923470",
	             "status": "PROCESSED",
	             "accrual": 500,
	             "uploaded_at": "2020-12-10T15:15:45+03:00"
	         },
	         {
	             "number": "12345678903",
	             "status": "PROCESSING",
	             "uploaded_at": "2020-12-10T15:12:01+03:00"
	         },
	         {
	             "number": "346436439",
	             "status": "INVALID",
	             "uploaded_at": "2020-12-09T16:09:53+03:00"
	         }
	     ]

	   204 — нет данных для ответа.
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

	orderHistory, err := h.Storage.SelectOrders(ctx, &user)
	if err != nil {
		fmt.Println("Проблемы с получением истории заказов", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	bodyBytes, err := json.Marshal(orderHistory)
	if err != nil {
		fmt.Println("Проблемы при маршалинге истории заказов", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(bodyBytes))
}
