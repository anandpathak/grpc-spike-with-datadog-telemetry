package telemetry

import (
	"log"

	"go.opencensus.io/stats/view"
)

type prometheusPushGateway struct{}

func (ppg *prometheusPushGateway) ExportView(vd *view.Data) {
	log.Printf("vd.View: %+v\n%#v\n", vd.View, vd.Rows)
	for i, row := range vd.Rows {
		log.Printf("\tRow: %#d: %#v\n", i, row)
	}
	log.Printf("StartTime: %s EndTime: %s\n\n", vd.Start.Round(0), vd.End.Round(0))
}

type Exporter interface {
	ExportView(vd *view.Data)
}
