package model

type UserForm struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

type SessionModel struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	Device     string `json:"device"`
	Created_At int64  `json:"created_at"`
	Last_Login *int64 `json:"last_login"`
	Expiry     int64  `json:"expiry"`
}
