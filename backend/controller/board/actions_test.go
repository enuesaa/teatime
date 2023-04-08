package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckin(t *testing.T) {
	c, response := contextForTest("/api/v1.Board/Checkin", `{"id":"aaa"}`)
	boardControllerForTest().Checkin(c)
	assert.Equal(t, 200, c.Writer.Status())
	assert.Equal(t, `{"id":"aaa"}`, response.Body.String())
}

func TestListTimeline(t *testing.T) {
	c, response := contextForTest("/api/v1.Board/ListTimeline", `{"id":"aaa","page":1}`)
	boardControllerForTest().ListTimeline(c)
	assert.Equal(t, 200, c.Writer.Status())
	assert.Equal(t, `{"id":"aaa","page":1,"items":[{"time":"2023-03-20T00:00:00+09:00"},{"time":"2023-03-21T00:00:00+09:00"}]}`, response.Body.String())
}

func TestArchive(t *testing.T) {
	c, response := contextForTest("/api/v1.Board/Archive", `{"id":"aaa"}`)
	boardControllerForTest().Delete(c)
	assert.Equal(t, 200, c.Writer.Status())
	assert.Equal(t, `{"id":"aaa"}`, response.Body.String())
}

func TestUnArchive(t *testing.T) {
	c, response := contextForTest("/api/v1.Board/UnArchive", `{"id":"aaa"}`)
	boardControllerForTest().Delete(c)
	assert.Equal(t, 200, c.Writer.Status())
	assert.Equal(t, `{"id":"aaa"}`, response.Body.String())
}
