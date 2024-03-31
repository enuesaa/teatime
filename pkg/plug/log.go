package plug

import (
	"fmt"

	"github.com/hashicorp/go-hclog"
)

func Log(format string, args ...any) {
	logger := hclog.New(&hclog.LoggerOptions{
		JSONFormat: true,
	})

	message := fmt.Sprintf(format, args...)
	logger.Info(message)
}

func LogE(err error) {
	logger := hclog.New(&hclog.LoggerOptions{
		JSONFormat: true,
	})

	message := fmt.Sprintf("Error: %s", err.Error())
	logger.Info(message)
}
