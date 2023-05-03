package model

type User struct {
	UserId int64  `json:"userid"`
	Name   string `json:"name"`
	Mobile int64  `json:"mobile"`
}
