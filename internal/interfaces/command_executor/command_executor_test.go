package command_executor

import (
	"log"
	"testing"
)

var (
	sampleOutput = "/root/api/flight-tracker/ws.sh####/root/api/flight-tracker####/root/api####GET /alpha/flight/api/ws/status HTTP/"
)

func TestParseOutput(t *testing.T) {
	out := ParseOutput(sampleOutput)
	log.Printf(out)
}
