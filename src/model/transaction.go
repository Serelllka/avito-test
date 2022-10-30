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
	Description string
	Amount      uint
	Type        TransactionType
}
