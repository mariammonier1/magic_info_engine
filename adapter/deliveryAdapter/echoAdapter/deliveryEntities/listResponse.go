package deliveryEntities

type ListResponse struct {
	Data   interface{} `json:"data"`
	Limit  int         `json:"limit"`
	Page   int         `json:"page"`
	Offset int         `json:"offset"`
	Count  int         `json:"count"`
}
