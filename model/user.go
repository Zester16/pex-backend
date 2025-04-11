package model

type UserForm struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

type SessionModel struct {
	Id         string
	Username   string
	Device     string
	Created_At int64
	Last_Login *int64
	Expiry     int64
}
