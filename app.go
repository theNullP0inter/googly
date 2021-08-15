package googly

import (
	"github.com/sarulabs/di/v2"
)

// App is any thing that Builds. build happens through inject its components into googlys builder
//
// Build(*di.Builder) is where you should inject dependencies for the sub-app
type App interface {
	Build(*di.Builder)
}
