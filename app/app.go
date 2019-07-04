package app

///go:generate mockgen -destination=testing.generated.go -package=app -source=app.go App
///go:generate mockgen -destination=testing.generated.go -package=app github.com/powerman/bug-gomock/app App
///go:generate mockgen -destination=testing.generated.go -package=app -self_package=github.com/powerman/bug-gomock/app -source=app.go App
///go:generate mockgen -destination=testing.generated.go -package=app -self_package=github.com/powerman/bug-gomock/app github.com/powerman/bug-gomock/app App

import (
	"context"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

type Nothing struct{}
type nothing struct{}

type App interface {
	// Noop(Ctx)             // source: doesn't work, reflect: works
	// Noop(context.Context) // source: work,         reflect: works
	// Noop(Nothing)         // source: doesn't work, reflect: doesn't work
	// Noop(nothing)         // source: work,         reflect: doesn't work
}
