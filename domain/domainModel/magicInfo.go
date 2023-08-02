package domainModel


type MagicInfoResponseMessage struct {
	Status       string `json:"status"`
	ErrorMessage error  `json:"errorMessage"`
	Items        interface{}   `json:"items"`
	TotalCount   int    `json:"totalCount"`
	PageSize     int    `json:"pageSize"`
	StartIndex   int    `json:"startIndex"`
}

type User struct {
	UserId       string `json:"userId"`
	UserName     string `json:"userName"`
	Email        string `json:"email"`
	Mobile       string `json:"mobile"`
	Phone        string `json:"phone"`
	Organization string `json:"organization"`
	GoupName     string `json:"goupName"`
	Role         string `json:"role"`
	Team         string `json:"team"`
	Position     string `json:"position"`
}
//
