package models

type User struct {
	IdUser   string `json:"id"`
	Name     string `json:"names"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
	Age      string `json:"age"`
}

type Color struct {
	Primary   string `json:"primary" bson:"primary"`
	Secondary string `json:"secondary" bson:"secondary"`
	Tertiary  string `json:"tertiary" bson:"tertiary"`
}
