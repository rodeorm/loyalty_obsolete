package main

import (
	"log"
	"loyalty/internal/api"
	"loyalty/internal/client"
	"loyalty/internal/repo"
)

/*
Система представляет собой HTTP API со следующими требованиями к бизнес-логике:
регистрация, аутентификация и авторизация пользователей;
приём номеров заказов от зарегистрированных пользователей;
учёт и ведение списка переданных номеров заказов зарегистрированного пользователя;
учёт и ведение накопительного счёта зарегистрированного пользователя;
проверка принятых номеров заказов через систему расчёта баллов лояльности;
начисление за каждый подходящий номер заказа положенного вознаграждения на счёт лояльности пользователя.
*/

func main() {
	databaseURI, runAddress, accrualSystemAddress := config()
	db, err := repo.InitPostgres(databaseURI)
	if err != nil {
		log.Fatal("ошибка при инициализации БД: ", err)
		return
	}
	go client.StartClient(db, accrualSystemAddress)
	api.StartAPI(db, runAddress)

}
