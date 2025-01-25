package plug

import (
	"fmt"

	"github.com/hashicorp/go-hclog"
)

type Logger struct {
	hclogger hclog.Logger
	teapod string
}

func (l *Logger) Info(format string, a... any) {
	text := fmt.Sprintf(format, a...)
	text = fmt.Sprintf("%s | %s", l.teapod, text)
	l.hclogger.Info(text)
}

func (l *Logger) Err(err error) {
	text := fmt.Sprintf("Error: %s", err.Error())
	text = fmt.Sprintf("%s | %s", l.teapod, text)
	l.hclogger.Info(text)
}
