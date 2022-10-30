package dto

type Transaction struct {
	ProducerId  int    `json:"producerId"`
	ConsumerId  int    `json:"consumerId"`
	Description string `json:"description"`
	Amount      uint   `json:"amount"`
}
