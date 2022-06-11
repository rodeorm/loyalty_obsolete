package model

//Пользователь
type User struct {
	Login     int // Пока просто число
	Password  string
	Balance   int64
	Withdrawn int64
}

//Заказ
type Order struct {
	User    *User
	Number  string
	Status  *OrderStatus
	Accrual int
}

//Операция над баллами лояльности
type Operation struct {
	Order *Order
	Sum   int64
	Type  *OperationType
}

/* Тип операции над баллами лояльности: списание/начисление. Для возможностей расширения в будущем:
реализовано списание по константе, для улучшения производительности и уменьшения нагрузки на I/O в СУБД - реализовать можно по ключу, а можно просто определять был
плюс или минус суммы в Operation*/
type OperationType struct {
	ID    int
	Name  string
	Const string
}

/* Статус операции над баллами лояльности Для возможностей расширения в будущем:
реализовано списание по константе, для улучшения производительности и уменьшения нагрузки на I/O в СУБД - реализовать можно по ключу */
type OrderStatus struct {
	ID    int
	Name  string
	Const string
}
