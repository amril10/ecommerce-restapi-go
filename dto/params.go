package dto

type ResponseWithParams struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
	Paginate   *Paginate
	Data       any
}