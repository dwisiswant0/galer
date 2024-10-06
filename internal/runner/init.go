package runner

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
)

func init() {
	clog = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
	})
}
