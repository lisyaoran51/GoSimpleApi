package dto

type Response struct {
	Success bool   `json:"success"`
	Reason  string `json:"reason"`
}
