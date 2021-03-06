package models

import (
	"time"
)

type User struct {
	Id                    int    `valid:"numeric"`
	Firstname             string `valid:"alphanum,required"`
	Lastname              string `valid:"alphanum,required"`
	Email                 string `valid:"email,required"`
	Password              string `valid:"alphanum,required"`
	Password_confirmation string `valid:"alphanum,required"`
	City                  string `valid:"alphanum"`
	State                 string `valid:"alphanum"`
	Country               string `valid:"alphanum"`
	User_thumbnail        string
	Mobile_number         string `valid:"numeric"`
	Devise_token          string `valid:"alphanum,required"`
	Status                bool
	Status_message        string
}

type SignUpSuccessMessage struct {
	Success string
	Message string
	User    User
}

type SignUpErrorMessage struct {
	Success string
	Error   string
}

type Session struct {
	SessionId int
	StartTime time.Time
}

type Login struct {
	Success string
	Message string
	User    User
	Session Session
}

type LogInSuccessMessage struct {
	Success string
	Message string
	User    User
}

type LogInErrorMessage struct {
	Success string
	Error   string
}

type LogOutSuccessMessage struct {
	Success string
	Message string
	User    User
}

type LogOutErrorMessage struct {
	Success string
	Error   string
}
