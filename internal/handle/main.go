package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// Deprecated
type Handle struct {
	c    *gin.Context
	data map[string]interface{}
}

type WithValidator interface {
	Validate() error
}

func Bind(c *gin.Context, body WithValidator) Handle {
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"erraor": err.Error()})
		c.Abort()
		return Handle{c: c, data: make(map[string]interface{}, 0)}
	}
	if err := body.Validate(); err != nil {
		c.JSON(400, gin.H{"erraaor": err.Error()})
		c.Abort()
		return Handle{c: c, data: make(map[string]interface{}, 0)}
	}
	return Handle{c: c, data: make(map[string]interface{}, 0)}
}

type Fetcher func() any

func (handle *Handle) Data(name string, fetcher Fetcher) {
	if handle.c.IsAborted() {
		return
	}
	handle.data[name] = fetcher()
}

type Processor func()

func (handle *Handle) Process(processor Processor) {
	if handle.c.IsAborted() {
		return
	}
	processor()
}

func (handle *Handle) Render(response interface{}) {
	if handle.c.IsAborted() {
		return
	}
	if val, ok := handle.data["*"]; ok {
		mapstructure.Decode(val, response)
	} else {
		mapstructure.Decode(handle.data, response)
	}

	handle.c.JSON(200, response)
}
