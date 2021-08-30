package entity

type Response struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}

