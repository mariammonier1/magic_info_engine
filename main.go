package main

import (
	"magic_info_engine/adapter/deliveryAdapter"
	messageDistributor "magic_info_engine/adapter/messagingAdapter"
	"magic_info_engine/adapter/messagingAdapter/producers"
	"magic_info_engine/domain"
	"magic_info_engine/domain/domainCommon/domainLogs"
	"magic_info_engine/port/outputPort"
	"time"

	"github.com/getsentry/sentry-go"
)

var outputPortInstance outputPort.IOutputPort

func init() {
	messagingProducerPortInstance := producers.NewNatsProducer()
	outputPortInstance = outputPort.IOutputPort{
		Producer:   messagingProducerPortInstance,
	}
}

func main() {

	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://c72ca03b2d5d4e71b29b8b430d0be232@o83459.ingest.sentry.io/5669820",
	})
	if err != nil {
		domainLogs.GetLogInstance().Error("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	domainAdapter := domain.NewDomain(outputPortInstance)
	domainServices := domain.NewDomain(outputPortInstance)
	_ = messageDistributor.NewConsumerAdapter(domainServices)
	_ = messageDistributor.NewResponderAdapter(domainServices)
	domainLogs.GetLogInstance().Debug("Service Started")
	deliveryAdapter.NewDeliveryAdapter(domainAdapter)
	domainLogs.GetLogInstance().Debug("Service Ended")

}
