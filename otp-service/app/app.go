package app

import (
	"go.uber.org/zap"
)

type AppContext struct {
	config       *Config
	logger       *zap.Logger
	Dependencies *Dependencies
}

func Init() *AppContext {
	appCtx := AppContext{}
	dep := NewDependencies()
	appCtx.Dependencies = dep
	LoadDependecies()
	return nil
}

func (appCtx *AppContext) Start() {
}

func (appCtx *AppContext) Stop() {

}
