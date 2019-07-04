package app

//go:generate mockgen -source=app.go -destination=testing.generated.go -package=app App

import (
	"context"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

// App provides application features service.
type App interface {
	Noop(ctx context.Context) // works
	// Noop(ctx Ctx)             // doesn't work
}
