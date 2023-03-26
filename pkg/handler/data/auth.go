package data

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserID string `json:"userId"`
	Token  string `json:"token"`
}
