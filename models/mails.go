package models

type (
	MailInfo struct {
		To       string `json:"to"`
		Token    string `json:"token"`
		Template string `json:"template"`
	}
)
