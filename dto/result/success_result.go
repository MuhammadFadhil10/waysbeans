package dto

type SuccessResult struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
