package dto

type Transaction struct {
	ProducerId  int    `json:"producerId"`
	ConsumerId  int    `json:"consumerId"`
	ServiceId   int    `json:"serviceId"`
	Description string `json:"description"`
	Amount      uint   `json:"amount"`
}
