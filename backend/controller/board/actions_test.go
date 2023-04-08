package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckin(t *testing.T) {
	c, _ := contextForTest("/api/v1.Board/Checkin", `{"id":"aaa"}`)
	boardControllerForTest().Delete(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestListTimeline(t *testing.T) {
	c, _ := contextForTest("/api/v1.Board/ListTimeline", `{"id":"aaa"}`)
	boardControllerForTest().Delete(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestArchive(t *testing.T) {
	c, _ := contextForTest("/api/v1.Board/Archive", `{"id":"aaa"}`)
	boardControllerForTest().Delete(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestUnArchive(t *testing.T) {
	c, _ := contextForTest("/api/v1.Board/UnArchive", `{"id":"aaa"}`)
	boardControllerForTest().Delete(c)
	assert.Equal(t, 200, c.Writer.Status())
}
