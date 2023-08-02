package deliveryCommon

import (
	"magic_info_engine/domain/domainCommon"

	"github.com/derekstavis/go-qs"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

type RequestQuery struct {
	Page    int               `json:"page" query:"page"`
	Limit   int               `json:"limit" query:"limit"`
	Sort    string            `json:"sort"`
	Search  string            `json:"search"`
	Filters map[string]Filter `json:"filters"`
	State   string            `json:"state"`
	Offset  int               `json:"offset" query:"offset"`
}

type Filter struct {
	Key      string `json:"key"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type CustomBinder struct{}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
	//You may use default binder
	db := new(echo.DefaultBinder)
	_ = db.Bind(i, c)
	if err = db.Bind(i, c); err != echo.ErrUnsupportedMediaType {
		// Define your custom implementation
		var query map[string]interface{}
		if query, err = qs.Unmarshal(c.QueryString()); err == nil {
			cfg := &mapstructure.DecoderConfig{
				Metadata: nil,
				Result:   &i,
				TagName:  "json",
			}
			decoder, _ := mapstructure.NewDecoder(cfg)
			_ = decoder.Decode(query)

			return
		}
	}
	return
}

func (o RequestQuery) MapToDomainModel() *domainCommon.QueryParameters {
	domainFilters := map[string]domainCommon.Filter{}
	for s, filter := range o.Filters {
		domainFilters[s] = domainCommon.Filter{
			Key:      filter.Key,
			Operator: filter.Operator,
			Value:    filter.Value,
		}
	}
	return &domainCommon.QueryParameters{
		Offset:  o.Offset,
		Limit:   o.Limit,
		Sort:    o.Sort,
		Search:  o.Search,
		Filters: domainFilters,
	}
}
