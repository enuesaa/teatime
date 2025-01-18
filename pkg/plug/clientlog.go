package plug

import (
	"encoding/json"
	"io"

	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/repository/db"
)

type ClientLogWriter struct {
	io.Writer
	repos repository.Repos
}

func (w *ClientLogWriter) Write(p []byte) (int, error) {
	type PluginLog struct {
		Created string `json:"timestamp"`
		Message string `json:"@message"`
	}
	query := w.repos.DB.Logs()

	var l PluginLog
	if err := json.Unmarshal(p, &l); err != nil {
		return 0, err
	}
	data := db.Log{
		Created: l.Created,
		Message: l.Message,
	}
	if _, err := query.Create(data); err != nil {
		return 0, err
	}
	return len(p), nil
}
