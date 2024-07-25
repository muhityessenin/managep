package model

type User struct {
	ID               string `json:"id"`
	FullName         string `json:"full_name"`
	Email            string `json:"email"`
	RegistrationDate string `json:"registration_date" datetime_format:"YYYY-MM-DD"`
	Role             string `json:"role"`
}
