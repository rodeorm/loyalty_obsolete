package api

import (
	"context"
	"io"
	"net/http"

	"github.com/rodeorm/loyalty/internal/api/cookie"
	"github.com/rodeorm/loyalty/internal/model"
)

/*
	  	Загрузка номера заказа
		Хендлер: POST /api/user/orders.
		Хендлер доступен только аутентифицированным пользователям. Номером заказа является последовательность цифр произвольной длины.
		Номер заказа может быть проверен на корректность ввода с помощью алгоритма Луна.
*/
func (h Handler) ordersPost(w http.ResponseWriter, r *http.Request) {
	/*
	   Формат запроса:
	   POST /api/user/orders HTTP/1.1
	   Content-Type: text/plain
	   ...

	   12345678903
	   Возможные коды ответа:
	   200 — номер заказа уже был загружен этим пользователем;
	   202 — новый номер заказа принят в обработку;
	   400 — неверный формат запроса;
	   401 — пользователь не аутентифицирован;
	   409 — номер заказа уже был загружен другим пользователем;
	   422 — неверный формат номера заказа;
	   500 — внутренняя ошибка сервера.
	*/
	userKey, err := cookie.GetUserKeyFromCoockie(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx := context.TODO()
	if r.ContentLength == 0 {
		w.WriteHeader(http.StatusBadRequest) //неверный формат запроса
		return
	}
	bodyBytes, _ := io.ReadAll(r.Body)
	numberFromBody := string(bodyBytes)
	if !model.CheckOrderNum(numberFromBody) {
		w.WriteHeader(http.StatusUnprocessableEntity) // неверный формат номера заказа
		return
	}
	order := model.Order{Number: numberFromBody, UserLogin: userKey, Status: "NEW"}
	dublicateOrder, anotherUserOrder, err := h.Storage.InsertOrderSimple(ctx, &order)
	switch {
	case err != nil: // внутренняя ошибка сервера
		w.WriteHeader(http.StatusInternalServerError)
	case anotherUserOrder: // номер заказа уже был загружен другим пользователем
		w.WriteHeader(http.StatusConflict)
	case dublicateOrder: // номер заказа уже был загружен этим пользователем
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusAccepted) //новый номер заказа принят в обработку
		return
	}

}
