package entity

import "time"

type User struct {
	UserId    uint64    `json:"userId"`
	Account   string    `json:"account"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}
