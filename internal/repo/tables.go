package repo

import (
	"context"
	"fmt"
)

func (s postgresStorage) createTables(ctx context.Context) error {

	err := s.dropTables(ctx)
	if err != nil {
		fmt.Println("Проблема при удалении таблиц")
		return err
	}

	err = s.createUsersTable(ctx)
	if err != nil {
		fmt.Println("Проблема при создании таблицы с Пользователями")
		return err
	}

	err = s.createOrdersTable(ctx)
	if err != nil {
		fmt.Println("Проблема при создании таблицы с Заказами")
		return err
	}

	err = s.createOperationsTable(ctx)
	if err != nil {
		fmt.Println("Проблема при создании таблицы с Операциями")
		return err
	}
	/*
		err = s.createOperationTypesTable(ctx)
		if err != nil {
			fmt.Println("Проблема при создании таблицы с Типами операции")
			return err
		}

		err = s.createOrderStatusesTable(ctx)
		if err != nil {
			fmt.Println("Проблема при создании таблицы со Статусами заказа")
			return err
		}
	*/
	return nil
}

func (s postgresStorage) dropTables(ctx context.Context) error {
	_, err := s.DB.ExecContext(ctx, "DROP TABLE IF EXISTS OperationTypes, OrderStatuses, Users, Orders, Operations")
	if err != nil {
		fmt.Println("Проблема при создании таблицы с Типами операции")
		return err
	}
	return nil
}

func (s postgresStorage) createUsersTable(ctx context.Context) error {
	_, err := s.DB.ExecContext(ctx,
		"CREATE TABLE IF NOT EXISTS  Users"+
			"("+
			"login VARCHAR(20) PRIMARY KEY"+
			", password VARCHAR(100) NOT NULL"+
			", balance MONEY"+
			");")
	if err != nil {
		fmt.Println("Проблема при создании таблицы с Пользователями: ", err)
		return err
	}
	return nil
}

func (s postgresStorage) createOrdersTable(ctx context.Context) error {
	_, err := s.DB.ExecContext(ctx,
		"CREATE TABLE IF NOT EXISTS Orders"+
			"("+
			"Number VARCHAR(20) PRIMARY KEY "+
			", Status VARCHAR(60) "+
			", UserLogin VARCHAR(20) REFERENCES Users(login) "+
			", Accrual MONEY "+
			", Withdrawn MONEY "+
			", UploadedTime TIMESTAMP "+
			", ProcessedTime TIMESTAMP "+
			");")
	if err != nil {
		fmt.Println("Проблема при создании таблицы с Заказами: ", err)
		return err
	}
	return nil
}

func (s postgresStorage) createOperationsTable(ctx context.Context) error {
	_, err := s.DB.ExecContext(ctx,
		"CREATE TABLE IF NOT EXISTS Operations"+
			"("+
			"ID SERIAL PRIMARY KEY"+
			", OrderNumber VARCHAR(20)"+
			", Sum MONEY"+
			", ProcessedTime TIMESTAMP "+
			", UserLogin VARCHAR(20) REFERENCES Users(login)"+
			")")
	if err != nil {
		fmt.Println("Проблема при создании таблицы с Операциями: ", err)
		return err
	}
	return nil
}

/*
func (s postgresStorage) createOperationTypesTable(ctx context.Context) error {
	_, err := s.DB.ExecContext(ctx,
		"CREATE TABLE IF NOT EXISTS  OperationTypes"+
			"("+
			"	ID INT PRIMARY KEY"+
			"	, Name VARCHAR(10) NOT NULL"+
			"	, Const VARCHAR(10) NOT NULL"+
			"); "+
			"INSERT INTO OperationTypes"+
			"("+
			"	Name"+
			"	, Const"+
			")"+
			"SELECT 1, 'Начисление', 'add' "+
			"UNION ALL "+
			"SELECT 2, 'Списание', 'remove';")

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s postgresStorage) createOrderStatusesTable(ctx context.Context) error {
	_, err := s.DB.ExecContext(ctx,
		"CREATE TABLE IF NOT EXISTS  OrderStatuses"+
			"("+
			"ID INT PRIMARY KEY"+
			", Name VARCHAR(150) NOT NULL"+
			", Const VARCHAR(20) NOT NULL"+
			");"+
			"INSERT INTO OrderStatuses"+
			"("+
			"Name"+
			", Const"+
			")"+
			"SELECT 1, 'Заказ загружен в систему, но не попал в обработку', 'NEW'"+
			" UNION ALL "+
			"SELECT 2, 'Вознаграждение за заказ рассчитывается', 'PROCESSING'"+
			" UNION ALL "+
			"SELECT 3, 'Cистема расчёта вознаграждений отказала в расчёте', 'INVALID'"+
			" UNION ALL "+
			"SELECT 4, 'Данные по заказу проверены, и информация о расчёте успешно получена', 'PROCESSED';")
	if err != nil {
		fmt.Println("Проблема при создании таблицы со Статусами заказов: ", err)
		return err
	}
	return nil
}
*/
