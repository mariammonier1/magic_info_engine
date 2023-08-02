package domainModel

import (
	"time"
)

type Component struct {
	Id        string      `json:"id"`
	Type      string      `json:"type"`
	IsLocked  bool        `json:"isLocked"`
	ProjectId string      `json:"projectId"`
	PageId    string      `json:"pageId"`
	Name      string      `json:"name"`
	Style     interface{} `json:"style"`
	Mapping   interface{} `json:"mapping"`
	Setting   interface{} `json:"settings"`
	Events    interface{} `json:"events"`
	Data      interface{} `json:"data"`
	Metadata  interface{} `json:"metadata"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
	ParentId  string      `json:"parentId"`
}

//
