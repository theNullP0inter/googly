package googly

import (
	"github.com/sarulabs/di/v2"
)

// AppInterface should be implemented by sub-apps
//
// Build(*di.Builder) is where you should inject dependencies for the sub-app
type AppInterface interface {
	Build(*di.Builder)
}
