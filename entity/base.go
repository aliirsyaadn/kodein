package entity

type Response struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}

var ResponseSuccess = Response{
	Message: "succes",
	Status: 200,
}

