package models

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

type SignUpMessage struct {
	Success string
	Message string
	User    User
}
