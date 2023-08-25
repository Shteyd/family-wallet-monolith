package customer

import "time"

type Customer struct {
	Id        int       `json:"id" example:"1"`
	Username  string    `json:"username" example:"Shteyd"`
	Email     string    `json:"email" example:"example@mail.com"`
	Password  string    `json:"password" example:"qwerty123"`
	CreatedAt time.Time `json:"created_at" example:""`
	UpdatedAt time.Time `json:"updated_at" example:""`
}
