package teletype

import (
	"context"
)

type Client interface {
	JSON(context.Context, string, interface{}) error
}

type Cleaner interface {
	Clean(string, ...string) (string, error)
}
