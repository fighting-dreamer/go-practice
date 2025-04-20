package handler

import "net/http"

type SystemHandler struct{

}

func NewSystemHandler() *SystemHandler {
	return &SystemHandler{}
}

func (sh *SystemHandler) Ping(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("pong"));
}