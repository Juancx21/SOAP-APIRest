package core

import "encoding/xml"

type Response struct {
	Error bool   `json:"error"`
	Data  User   `json:"data"`
	Msg   string `json:"msg"`
	Code  int    `json:"code"`
	Type  string `json:"type"`
}

type User struct {
	IdUser   string `json:"id"`
	Name     string `json:"names"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
	Age      string `json:"age"`
}

type RequestUserById struct {
	IdUser string `json:"id"`
}

type RequestUserByEmail struct {
	Email string `json:"email"`
}

type RequestDeleteUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestCreateUser struct {
	IdUser   string `json:"id"`
	Name     string `json:"names"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
	Age      string `json:"age"`
	Gender   string `json:"gender"`
	Password string `json:"password"`
}

type ResponseByID struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	SOAPENV string   `xml:"SOAP-ENV,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text                string `xml:",chardata"`
		GetUserByIdResponse struct {
			Text string `xml:",chardata"`
			Ns2  string `xml:"ns2,attr"`
			User struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id"`
				Names    string `xml:"names"`
				Lastname string `xml:"lastname"`
				Email    string `xml:"email"`
				Age      string `xml:"age"`
				Password string `xml:"password"`
			} `xml:"user"`
		} `xml:"getUserByIdResponse"`
	} `xml:"Body"`
}

type ResponseByEmail struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	SOAPENV string   `xml:"SOAP-ENV,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text                   string `xml:",chardata"`
		GetUserByEmailResponse struct {
			Text string `xml:",chardata"`
			Ns2  string `xml:"ns2,attr"`
			User struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id"`
				Names    string `xml:"names"`
				Lastname string `xml:"lastname"`
				Email    string `xml:"email"`
				Age      string `xml:"age"`
				Password string `xml:"password"`
			} `xml:"user"`
		} `xml:"getUserByEmailResponse"`
	} `xml:"Body"`
}

type ResponseCreateUser struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	SOAPENV string   `xml:"SOAP-ENV,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text                string `xml:",chardata"`
		GetUserByIdResponse struct {
			Text string `xml:",chardata"`
			Ns2  string `xml:"ns2,attr"`
			User struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id"`
				Names    string `xml:"names"`
				Lastname string `xml:"lastname"`
				Email    string `xml:"email"`
				Age      string `xml:"age"`
				Gender   string `xml:"gender"`
			} `xml:"user"`
		} `xml:"getUserByIdResponse"`
	} `xml:"Body"`
}

type ResponseDeleteUser struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	SOAPENV string   `xml:"SOAP-ENV,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text               string `xml:",chardata"`
		DeleteUserResponse struct {
			Text    string `xml:",chardata"`
			Ns2     string `xml:"ns2,attr"`
			Message string `xml:"message"`
		} `xml:"deleteUserResponse"`
	} `xml:"Body"`
}
