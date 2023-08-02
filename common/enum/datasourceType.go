package enum

type DatasourceType string

const (
	Rest      DatasourceType = "REST"
	Cervello  DatasourceType = "CERVELLO"
	Postgres  DatasourceType = "POSTGRES"
	WebSocket DatasourceType = "WEB_SOCKET"
)
