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
Запрос на списание средств
Хендлер: POST /api/user/balance/withdraw
Хендлер доступен только авторизованному пользователю. Номер заказа представляет собой гипотетический номер нового заказа пользователя, в счёт оплаты которого списываются баллы.
*/
func (h Handler) balanceWithdrawPost(w http.ResponseWriter, r *http.Request) {
	/*
			Примечание: для успешного списания достаточно успешной регистрации запроса, никаких внешних систем начисления не предусмотрено и не требуется реализовывать.
			Формат запроса:
			POST /api/user/alance/withdraw HTTP/1.1
			Content-Type: application/json

		{
			   "order": "2377225624",
			    "sum": 751
			}
			Зесь order — номер заказа, а sum — сумма баллов к списанию в счёт оплаты.
			Возможные коды ответа:
			200 — успешная обработа запроса;
			401 — пользователь не авторизован
			402 — на счету недостаточно средст;
			422 — неверный номер заказа;
			500 — внутренняя ошибка сервра.
	*/
	userKey, err := cookie.GetUserKeyFromCoockie(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user := model.User{Login: userKey}
	ctx := context.TODO()

	h.Storage.SelectUserData(ctx, &user)
	operation := model.Operation{}
	bodyBytes, _ := ioutil.ReadAll(r.Body)

	err = json.Unmarshal(bodyBytes, &operation)
	operation.UserLogin = userKey
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.Balance < operation.Sum {
		w.WriteHeader(http.StatusPaymentRequired) //402 — на счету недостаточно средст
		return
	}

	if !model.CheckOrderNum(operation.OrderNumber) {
		w.WriteHeader(http.StatusUnprocessableEntity) // 422 — неверный номер заказа
		return
	}
	err = h.Storage.InsertOperation(ctx, &operation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
