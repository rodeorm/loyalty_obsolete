package api

/*
Накопительная система лояльности «Гофермарт» должна предоставлять следующие HTTP-хендлеры:
POST /api/user/register — регистрация пользователя;
POST /api/user/login — аутентификация пользователя;
POST /api/user/orders — загрузка пользователем номера заказа для расчёта;
GET /api/user/orders — получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях;
GET /api/user/balance — получение текущего баланса счёта баллов лояльности пользователя;
POST /api/user/balance/withdraw — запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа;
GET /api/user/balance/withdrawals — получение информации о выводе средств с накопительного счёта пользователем.
*/

import (
	"log"
	"loyalty/internal/repo"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//RouterStart запускает веб-сервер
func StartAPI(storage repo.Storage, runAddress, accrualSystemAddress string) error {

	r := mux.NewRouter()
	h := Handler{Storage: storage}

	//POST /api/user/register — регистрация пользователя
	r.HandleFunc("/api/user/register", h.registerPost).Methods(http.MethodPost)

	//POST /api/user/login — аутентификация пользователя
	r.HandleFunc("/api/user/login", h.loginPost).Methods(http.MethodPost)

	//POST /api/user/orders — загрузка пользователем номера заказа для расчёта
	r.HandleFunc("/api/user/orders", h.ordersPost).Methods(http.MethodPost)

	//GET /api/user/orders — получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях
	r.HandleFunc("/api/user/orders", h.ordersGet).Methods(http.MethodGet)

	//GET /api/user/balance — получение текущего баланса счёта баллов лояльности пользователя
	r.HandleFunc("/api/user/balance", h.balanceGet).Methods(http.MethodGet)

	//POST /api/user/balance/withdraw — запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа
	r.HandleFunc("/api/user/balance/withdraw", h.balanceWithdrawPost).Methods(http.MethodPost)

	//GET /api/user/balance/withdrawals — получение информации о выводе средств с накопительного счёта пользователем
	r.HandleFunc("/api/user/balance/withdrawals", h.balanceWithdrawalsGet).Methods(http.MethodGet)

	srv := &http.Server{
		Handler:      r,
		Addr:         runAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

	return nil
}

type Handler struct {
	Storage repo.Storage
}
