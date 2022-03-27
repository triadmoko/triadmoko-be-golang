package entity

import "time"

type Nakes struct {
	ID      int
	NoNakes int
	Name    string
	Addres  string
	FaskesID int
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time
}
