package types


type Student struct {
	Id string `json:"id"`
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age int `json:"age" validate:"required,min=0"`
	Password string `json:"password" validate:"required,min=6"`
}