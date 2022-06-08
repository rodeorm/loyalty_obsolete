package main

import (
	"flag"
	"os"
)

/*
Сервис должен поддерживать конфигурирование следующими методами:
адрес и порт запуска сервиса: переменная окружения ОС RUN_ADDRESS или флаг -a;
адрес подключения к базе данных: переменная окружения ОС DATABASE_URI или флаг -d;
адрес системы расчёта начислений: переменная окружения ОС ACCRUAL_SYSTEM_ADDRESS или флаг -r
*/
func config() (string, string, string) {
	flag.Parse()

	// os.Setenv("RUN_ADDRESS", "localhost:8080")
	os.Setenv("DATABASE_URI", "postgres://app:qqqQQQ123@localhost:5432/loyalty?sslmode=disable")
	// os.Setenv("ACCRUAL_SYSTEM_ADDRESS", "localhost:8080")

	var runAddress, databaseURI, accrualSystemAddress string

	//Адрес и порт запуска сервиса
	runAddress = *a
	if runAddress == "" {
		runAddress = os.Getenv("RUN_ADDRESS")
	}

	databaseURI = *d
	if databaseURI == "" {
		databaseURI = os.Getenv("DATABASE_URI")
	}

	accrualSystemAddress = *r
	if accrualSystemAddress == "" {
		accrualSystemAddress = os.Getenv("ACCRUAL_SYSTEM_ADDRESS")
	}

	return databaseURI, runAddress, accrualSystemAddress
}
