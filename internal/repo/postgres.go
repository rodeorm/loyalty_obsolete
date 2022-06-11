package repo

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type postgresStorage struct {
	DB *sql.DB // Драйвер подключения к СУБД
}

//InitPostgres cоздает хранилище данных в БД на экземпляре Postgres
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

func (s postgresStorage) CloseConnection() {
	s.DB.Close()
}
