package web

import "testing"
import "github.com/stretchr/testify/assert"

func TestRenderIndex(t *testing.T) {
	// TODO: Expand to actually perform tests
	web := Web{Addr: ":8080"}

	assert.NotNil(t, web)
}
