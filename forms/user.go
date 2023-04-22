package forms

type UserCreateForm struct {
	Email     string `json:"email" binding:"required,email,unique_email"`
	FirstName string `json:"first_name" binding:"required,alpha,min=2,max=32"`
	LastName  string `json:"last_name" binding:"required,alpha,min=2,max=32"`
}

type UserUpdateForm struct {
	FirstName string `json:"first_name" binding:"omitempty,alpha,min=2,max=32"`
	LastName  string `json:"last_name" binding:"omitempty,alpha,min=2,max=32"`
}
