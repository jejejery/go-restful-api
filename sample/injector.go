//go:build wireinject
// +build wireinject

package sample

import (
	"github.com/google/wire"
)

func InitializeSimpleService() *SimpleService {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return &SimpleService{}
}

