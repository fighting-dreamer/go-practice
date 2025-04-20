package app

import (
	"nipun.io/otp-service/server"
	"nipun.io/otp-service/service"
)

type Dependencies struct {
	Config *Config
	otpService service.IOTPService
	server server.IServer
}

func NewDependencies() *Dependencies {
	return &Dependencies{}
}

func LoadDependecies() {
	config:= LoadConfig()
	dependencies := NewDependencies()
	// add config
}

func (dep *Dependencies) addServer() {
	dep.server = server.NewHttpServer(dep.Config.Port)
}

func (dep *Dependencies) addServices() {
	dep.otpService = service.NewOTPService()
}

func 