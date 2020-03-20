package gotest

import (
	"time"
)

type User struct {
	ID string `json:"id"`
	Profile Profile `json:"profile"`
	PersonalInfo PersonalInfo `json:"personal_info"`
	CreateAt time.Time `json:"create_at"`
}


type Profile struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type PersonalInfo struct {
	Name string `json:"name"`
	LastName string `json:"last_name"`
	Address string `json:"address"`
	Age string `json:"age"`
	DNI string `json:"dni"`
}


