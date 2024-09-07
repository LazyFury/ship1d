package httpserver

import (
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"lazyfury.github.com/yoomall-server/config"
)

type HttpServer struct {
	Engine *gin.Engine
	Config *config.Config
}

func (h *HttpServer) Start() *gin.Engine {
	port := strconv.Itoa(h.Config.HTTP.Port)
	if port == "0" {
		port = "8900"
	}
	log.Info("[server on] http://127.0.0.1:" + port)
	h.Engine.Run(":" + port)
	return h.Engine
}
