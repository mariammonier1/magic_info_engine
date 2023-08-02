package natsAdapter

import (
	"magic_info_engine/common"
	"magic_info_engine/domain/domainCommon/domainLogs"
	"time"

	"github.com/nats-io/nats.go"
)

var config = common.GetEnvConfig()

func NewConnection() (*nats.EncodedConn, error) {
	mdConn, err := nats.Connect(
		config.MessagingHost,
		nats.UserInfo(config.MessagingUsername, config.MessagingPassword),
		nats.ReconnectWait(time.Second*2),
		nats.ReconnectHandler(func(_ *nats.Conn) {
			domainLogs.GetLogInstance().Debug("nats reconnecting")
		}),
		nats.DisconnectErrHandler(func(_ *nats.Conn, err error) {
			domainLogs.GetLogInstance().Error("nats disconnect", err)
		}))
	if err != nil {
		return nil, err
	}
	connection, err := nats.NewEncodedConn(mdConn, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}
	return connection, nil
}
