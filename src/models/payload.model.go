package models

type LogInPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
