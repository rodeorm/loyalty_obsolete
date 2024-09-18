package repo

import (
	"context"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type postgresStorage struct {
	DB *sqlx.DB // Драйвер подключения к СУБД
}

// InitPostgres cоздает хранилище данных в БД на экземпляре Postgres
func InitPostgres(connectionString string) (*postgresStorage, error) {
	db, err := sqlx.Open("pgx", connectionString)
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
