package models


type Userdata struct {
	ID 			uint	`json:"id"`
	Name		string	`json:"name"`
	Age			uint	`json:"age"`
	Sex 		string	`json:"sex"`
	SexChoose	string	`json:"sex_choose"`
	PhoneNum	string	`json:"phone_num"`
	Password	string	`json:"-"`

}