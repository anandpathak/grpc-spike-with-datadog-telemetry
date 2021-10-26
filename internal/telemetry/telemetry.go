package telemetry

import (
	"go.opencensus.io/plugin/ocgrpc"
	"log"

	datadog "github.com/DataDog/opencensus-go-exporter-datadog"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
)

func Telemetry() {

	if err := view.Register(ocgrpc.DefaultClientViews...); err != nil {
		log.Fatalf("Failed to register ocgrpc client views: %v", err)
	}
	dd, err := datadog.NewExporter(
		datadog.Options{
			Namespace:              "demogrpc",
			Service:                "demogrpc",
			TraceAddr:              "localhost:8125",
			StatsAddr:              "localhost:8125",
			GlobalTags:             nil,
			DisableCountPerBuckets: false,
			StatsdOptions:          nil,
		},
	)

	view.RegisterExporter(new(prometheusPushGateway))
	if err != nil {
		log.Fatalf("Failed to create the Datadog exporter: %v", err)
	}
	// It is imperative to invoke flush before your main function exits

	// Register it as a metrics exporter
	view.RegisterExporter(dd)

	// Register it as a metrics exporter

	// Allow Datadog to calculate APM metrics and do the sampling.
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.NeverSample()})
}
