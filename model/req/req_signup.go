package req

type RequestSignUp struct {
	FullName string `json:"fullName" validate:"required"`
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
