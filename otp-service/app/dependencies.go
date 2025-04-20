package app

import (
	"nipun.io/otp-service/server"
	"nipun.io/otp-service/service"
)

type Dependencies struct {
	Config     *Config
	otpService service.IOTPService
	server     server.IServer
}

func NewDependencies() *Dependencies {
	return &Dependencies{}
}

func LoadDependecies() {
	dependencies := NewDependencies()
	config := LoadConfig()
	dependencies.Config = config
	dependencies.addServices()
	dependencies.addServer()
}

func (dep *Dependencies) addServer() {
	// dep.server = server.NewHttpServer(dep.Config.Port)
}

func (dep *Dependencies) addServices() {
	// dep.otpService = service.NewOTPService()
}
