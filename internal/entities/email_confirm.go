package entities

import "fmt"

type EmailConfirm struct {
	Email             string
	EmailConfirmToken string
	Url               string
}

func NewEmailConfirm(email string, emailConfirmToken string, host string, path string) *EmailConfirm {
	url := fmt.Sprintf("%s%s?token=%s", host, path, emailConfirmToken)
	return &EmailConfirm{
		Email:             email,
		EmailConfirmToken: emailConfirmToken,
		Url:               url,
	}
}
