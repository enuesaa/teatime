package apptest

import (
	"net/http/httptest"

	"github.com/thedevsaddam/gojsonq/v2"
)

type Result struct {
	*httptest.ResponseRecorder
}

func (r *Result) ParseBody() *gojsonq.JSONQ {
	return gojsonq.New().FromString(r.Body.String())
}

func (r *Result) GetS(path string) string {
	return r.ParseBody().Find(path).(string)
}

func (r *Result) GetI(path string) int {
	return r.ParseBody().Find(path).(int)
}
