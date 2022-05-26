package entity

type RegisterRequest struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Health bool   `json:"health"`
}

type RegisterResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
