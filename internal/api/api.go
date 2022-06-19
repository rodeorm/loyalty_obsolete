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
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"loyalty/internal/api/middleware"
	"loyalty/internal/repo"
)

//RouterStart запускает веб-сервер
func StartAPI(storage repo.Storage, runAddress string) error {
	h := Handler{Storage: storage}

	// Те, кто НЕ попадают под middleware проверку аутентификации
	router := mux.NewRouter()
	// Те, кто попадают под middleware проверку аутентификации
	api := router.PathPrefix("/").Subrouter()
	api.Use(middleware.AuthMiddleware)

	// Регистрация пользователя
	router.HandleFunc("/api/user/register", h.registerPost).Methods(http.MethodPost)
	// Аутентификация пользователя
	router.HandleFunc("/api/user/login", h.loginPost).Methods(http.MethodPost)
	// При попытке доступа к методам без пройденной аутентификации
	router.HandleFunc("/forbidden", h.forbidden)

	//Загрузка пользователем номера заказа для расчёта
	api.HandleFunc("/api/user/orders", h.ordersPost).Methods(http.MethodPost)
	//Получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях
	api.HandleFunc("/api/user/orders", h.ordersGet).Methods(http.MethodGet)
	//Получение текущего баланса счёта баллов лояльности пользователя
	api.HandleFunc("/api/user/balance", h.balanceGet).Methods(http.MethodGet)
	//Запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа
	api.HandleFunc("/api/user/balance/withdraw", h.balanceWithdrawPost).Methods(http.MethodPost)
	//Получение информации о выводе средств с накопительного счёта пользователем
	api.HandleFunc("/api/user/withdrawals", h.balanceWithdrawalsGet).Methods(http.MethodGet)

	srv := &http.Server{
		Handler:      router,
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
