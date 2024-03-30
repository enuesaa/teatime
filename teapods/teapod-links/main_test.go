package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFmtKey(t *testing.T) {
	handler := Handler{}
	assert.Equal(t, "id:name", handler.fmtKey("id", "name"))
	assert.Equal(t, "aaaaa:bb", handler.fmtKey("aaaaa", "bb"))
}

func TestFmtPattern(t *testing.T) {
	handler := Handler{}
	assert.Equal(t, "id:*", handler.fmtPattern("id"))
	assert.Equal(t, "aaaaa:*", handler.fmtPattern("aaaaa"))
}
