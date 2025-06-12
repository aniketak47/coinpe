package controllers

type CreateAccountRequest struct {
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Role        string `json:"role" validate:"required"`
	IsPartial   bool   `json:"is_partial,omitempty"`
}
