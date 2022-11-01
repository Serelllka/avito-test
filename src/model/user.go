package model

type UserAccount struct {
	Id   uint
	Name string
}

type UserAccountBalance struct {
	Id       uint `db:"id"`
	Income   uint `db:"income"`
	Outcome  uint `db:"outcome"`
	Reserved uint `db:"reserved"`
}
