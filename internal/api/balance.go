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
Хендлер: GET /api/user/balance.
Хендлер доступен только авторизованному пользователю. В ответе должны содержаться данные о текущей сумме баллов лояльности, а также сумме использованных за весь период регистрации баллов.
*/
func (h Handler) balanceGet(w http.ResponseWriter, r *http.Request) {
	/*
		Хендлер: GET /api/user/balance.
		Хендлер доступен только авторизованному пользователю. В ответе должны содержаться данные
		о текущей сумме баллов лояльности, а также сумме использованных за весь период регистрации баллов.
	*/
	userKey, err := cookie.GetUserKeyFromCoockie(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx := context.TODO()
	//	fmt.Println("Возвращаем баланс для пользователя: ", userKey)
	user := model.User{Login: userKey}
	h.Storage.SelectUserData(ctx, &user)
	balance := model.Balance{Current: user.Balance, Withdrawn: user.Withdrawn}
	fmt.Println("Баланс равен:", balance)
	bodyBytes, err := json.Marshal(balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(bodyBytes))
}
