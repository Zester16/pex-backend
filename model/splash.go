package model

type Splash struct {
	Name string `json:"name"`

	Date int64 `json:"date"`
}

type SplashModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	Date int64 `json:"date"`
}
