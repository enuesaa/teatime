package ctx

import (
	"encoding/json"

	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context

	Repos repository.Repos
}

func (cc *Context) WithData(data interface{}) error {
	cc.Set("data", data)

	return nil
}

func (cc *Context) BindBody(data interface{}) error {
	reqbody := cc.Request().Body

	return json.NewDecoder(reqbody).Decode(data)
}
