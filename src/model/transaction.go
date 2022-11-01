package model

type TransactionType int

const (
	Deposit TransactionType = iota
	Remittance
	Reservation
	Payment
)

type Transaction struct {
	Id          int
	ProducerId  int
	ConsumerId  int
	ServiceId   int
	OrderId     int
	Description string
	Amount      uint
	Type        TransactionType
}
