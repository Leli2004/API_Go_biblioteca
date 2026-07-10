package entity

type RespError struct {
	Code int    `json:"code"`
	Err  string `json:"error"`
}
