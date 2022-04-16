package auth

type Auth struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
