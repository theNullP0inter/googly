package googly

import (
	"github.com/sarulabs/di/v2"
)

type AppInterface interface {
	Build(builder *di.Builder)
}
