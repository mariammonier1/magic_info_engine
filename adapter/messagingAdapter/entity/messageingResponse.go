package entity

type MessagingResponse struct {
	Result interface{} `json:"result"`
	Reply  interface{} `json:"reply"`
	Status string      `json:"status"`
}
