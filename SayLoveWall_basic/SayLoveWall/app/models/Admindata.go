package models

type Admindata struct {
	ID       uint   `json:"id"`
	Account  string `json:"Account"`
	Password string `json:"-"`
}
