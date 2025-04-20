package server

import (
	"github.com/gin-gonic/gin"
	"nipun.io/otp-service/server/handler"
)

type IServer interface {
	Start()
	Stop()
}

type HttpServer struct {
	Port string
	httpServer *gin.Engine
	systemHandler handler.SystemHandler 
	otpHandler handler.OTPHandler
}

func NewHttpServer(port string) *HttpServer {
	return &HttpServer{
		Port: port,
		httpServer: gin.Default(),
	}
}

func (s *HttpServer) Start(){
	
}

func (s *HttpServer) Stop(){
	
}