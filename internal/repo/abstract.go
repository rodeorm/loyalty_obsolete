package repo

import "log"

type Storage interface {
}

// NewStorage определяет место для хранения данных
func NewStorage(connectionString string) Storage {
	var storage Storage
	storage, err := InitPostgres(connectionString)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return storage
}
