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

func (s postgresStorage) createOperationTypesTable(ctx context.Context) error {
	_, err := s.DB.ExecContext(ctx,
		"CREATE TABLE IF NOT EXISTS  OperationTypes"+
			"("+
			"	ID SERIAL PRIMARY KEY"+
			"	, Name VARCHAR(10) NOT NULL"+
			"	, Const VARCHAR(10) NOT NULL"+
			"); "+
			"INSERT INTO OperationTypes"+
			"("+
			"	Name"+
			"	, Const"+
			")"+
			"SELECT 'Начисление', 'add' "+
			"UNION ALL "+
			"SELECT 'Списание', 'remove';")

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
			"ID SERIAL PRIMARY KEY"+
			", Name VARCHAR(150) NOT NULL"+
			", Const VARCHAR(20) NOT NULL"+
			");"+
			"INSERT INTO OrderStatuses"+
			"("+
			"Name"+
			", Const"+
			")"+
			"SELECT 'Заказ загружен в систему, но не попал в обработку', 'NEW'"+
			" UNION ALL "+
			"SELECT 'Вознаграждение за заказ рассчитывается', 'PROCESSING'"+
			" UNION ALL "+
			"SELECT 'Cистема расчёта вознаграждений отказала в расчёте', 'INVALID'"+
			" UNION ALL "+
			"SELECT 'Данные по заказу проверены, и информация о расчёте успешно получена', 'PROCESSED';")
	if err != nil {
		fmt.Println("Проблема при создании таблицы со Статусами заказов: ", err)
		return err
	}
	return nil
}

func (s postgresStorage) createUsersTable(ctx context.Context) error {
	_, err := s.DB.ExecContext(ctx,
		"CREATE TABLE IF NOT EXISTS  Users"+
			"("+
			"login INT PRIMARY KEY"+
			", password VARCHAR(10) NOT NULL"+
			", balance int"+
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
			"ID INT PRIMARY KEY "+
			", OrderStatusID INT REFERENCES OrderStatuses(ID) "+
			", UserLogin INT REFERENCES Users(login) "+
			", Accrual INT "+
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
			", OrderID INT REFERENCES Orders(ID)"+
			", OperationTypeID INT  REFERENCES OperationTypes(ID)"+
			", Sum INT"+
			")")
	if err != nil {
		fmt.Println("Проблема при создании таблицы с Операциями: ", err)
		return err
	}
	return nil
}
