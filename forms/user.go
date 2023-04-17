package forms

type UserCreateForm struct {
	Username string `json:"username" binding:"required,alphanum,min=4,max=255"`
	Email    string `json:"email" binding:"required,email,unique_email"`
	Password string `json:"password" binding:"required,min=8,max=255"`
}
