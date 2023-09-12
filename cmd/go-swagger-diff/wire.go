//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	// "github/jacktrane/go-swagger-diff/internal/biz"
	// "github/jacktrane/go-swagger-diff/internal/conf"

	"github.com/go-kratos/kratos/v2"
	// "github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp() (*kratos.App, func(), error) {
	panic(wire.Build(newApp))
}
