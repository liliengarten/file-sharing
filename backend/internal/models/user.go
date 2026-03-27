package models

type User struct {
	ID 	  string `json:"id"`
	FirstName string `json:"first_name" validate:"required,min=3",max=100"`
	LastName  string `json:"last_name"  validate:"required,min=3",max=100"`
	Username  string `json:"username"   validate:"required,min=6,max=50"`
	Email     string `json:"email"      validate:"required,email"`
	Password  string `json:"password"   validate:"required,min=6"`
}
