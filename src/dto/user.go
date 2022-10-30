package dto

type UserAccount struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateUser struct {
	Name string `json:"name"`
}
