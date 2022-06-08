package repo

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type postgresStorage struct {
	DB *sql.DB // Драйвер подключения к СУБД
}

//cоздает хранилище данных в БД на экземпляре Postgres
func InitPostgres(connectionString string) (*postgresStorage, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	storage := postgresStorage{DB: db}
	storage.createTables(ctx)

	return &storage, nil
}

func (s postgresStorage) createTables(ctx context.Context) error {
	_, err := s.DB.ExecContext(ctx,
		"DROP TABLE IF EXISTS Users;"+
			"CREATE TABLE IF NOT EXISTS  Users"+
			"("+
			"ID SERIAL PRIMARY KEY"+
			", Name VARCHAR(10) NULL"+
			", Password VARCHAR(10) NOT NULL"+
			")")
	if err != nil {
		fmt.Println("Проблема при создании таблиц")
		return err
	}
	return nil
}
