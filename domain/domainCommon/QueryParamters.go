package domainCommon

type QueryParameters struct {
	Page    int
	Limit   int
	Sort    string
	Search  string
	Filters map[string]Filter
	Offset  int
}

type Filter struct {
	Key      string `json:"key"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}
