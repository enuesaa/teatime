package plug

import (
	"fmt"

	"github.com/hashicorp/go-hclog"
)

type Logger struct {
	hclogger hclog.Logger
}

func (l *Logger) Info(format string, a... any) {
	text := fmt.Sprintf(format, a...)
	l.hclogger.Info(text)
}

func (l *Logger) Err(err error) {
	text := fmt.Sprintf("Error: %s", err.Error())
	l.hclogger.Info(text)
}
