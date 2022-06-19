package model

import "time"

//Пользователь
type User struct {
	Login     string `json:"login,omitempty"`
	Password  string `json:"password,omitempty"`
	Balance   int
	Withdrawn int
}

//Заказ
type Order struct {
	UserLogin      string    `json:"-"`
	Number         string    `json:"number,omitempty"`
	Status         string    `json:"status,omitempty"`
	AccrualBalls   int       `json:"accrual,omitempty"` //Рассчитанные баллы за заказ
	WithdrawnBalls int       `json:"-"`                 //Сумма баллов к списанию в счёт оплаты заказа
	UploadedTime   time.Time `json:"uploaded_at,omitempty"`
	ProcessedTime  time.Time `json:"-"`
}

//Баланс
type Balance struct {
	Current   int `json:"current,omitempty"`
	Withdrawn int `json:"withdrawn,omitempty"`
}

//Операция над баллами лояльности
type Operation struct {
	OrderNumber string
	Sum         int
}

//Заказ пользователя

/*





Тип операции над баллами лояльности: списание/начисление. Для возможностей расширения в будущем:
реализовано списание по константе, для улучшения производительности и уменьшения нагрузки на I/O в СУБД - реализовать можно по ключу, а можно просто определять был
плюс или минус суммы в Operation
type OperationType struct {
	ID    int
	Name  string
	Const string
}
Статус операции над баллами лояльности Для возможностей расширения в будущем:
реализовано списание по константе, для улучшения производительности и уменьшения нагрузки на I/O в СУБД - реализовать можно по ключу
type OrderStatus struct {
	ID    int
	Name  string
	Const string
}

type Session struct {
	User           *User
	ExpirationTime time.Time
}

type SessionStore struct {
	Sessions map[string]*Session
}
*/
