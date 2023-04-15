package data

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponse struct {
	UserID string `json:"userId"`
	Token  string `json:"token"`
}
