package validators

type CreateUser struct {
	Email string `json:"email" binding:"required,email"`
	User  string `json:"user" binding:"required"`
	Pass  string `json:"pass" binding:"required"`
}
