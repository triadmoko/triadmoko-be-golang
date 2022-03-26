package entity

import "time"

type User struct {
	ID        uint32 `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     uint64 `json:"phone"`
	Address   string `json:"address"`
	Role      uint8
	Image     string `json:"image"`
	Password  string
	CreateAt  time.Time
	UpdateAt  time.Time
	DeleteAt  time.Time
}
