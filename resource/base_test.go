package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theNullP0inter/googly/logger"
)

// TestBaseResourceManager creates a new BaseResourceManager
// using NewBaseResourceManager() and tests its methods
func TestBaseResourceManager(t *testing.T) {
	r := new(MockResource)
	l := logger.NewGooglyLogger()
	rm := NewBaseResourceManager(l, r)
	assert.Same(t, r, rm.GetResource())
}
