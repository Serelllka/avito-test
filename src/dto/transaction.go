package dto

type Transaction struct {
	ProducerId  int    `json:"producerId"`
	ConsumerId  int    `json:"consumerId"`
	ServiceId   int    `json:"serviceId"`
	OrderId     int    `json:"orderId"`
	Description string `json:"description"`
	Amount      uint   `json:"amount"`
}

type Deposit struct {
	ConsumerId  uint   `json:"consumerId" binding:"required"`
	Description string `json:"description" binding:"required"`
	Amount      uint   `json:"amount" binding:"required"`
}

type Remittance struct {
	ProducerId  uint   `json:"producerId" binding:"required"`
	ConsumerId  uint   `json:"consumerId" binding:"required"`
	Description string `json:"description" binding:"required"`
	Amount      uint   `json:"amount" binding:"required"`
}

type Reservation struct {
	ProducerId  int    `json:"producerId" binding:"required"`
	ServiceId   int    `json:"serviceId" binding:"required"`
	OrderId     int    `json:"orderId" binding:"required"`
	Description string `json:"description" binding:"required"`
	Amount      uint   `json:"amount" binding:"required"`
}
