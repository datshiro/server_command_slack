package handlers

import (
	"path"
	"server/internal/infras/server/handlers/slack_handler"
	"server/internal/interfaces"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(g *gin.Engine, apiPrefix string) {
	register(g, apiPrefix, slack_handler.NewHandler())
}

func register(g *gin.Engine, apiPrefix string, h interfaces.Handler) {
	path := path.Join(apiPrefix, h.GetPath())
	switch h.GetMethod() {
	case "GET":
		g.GET(path, h.Handle)
	case "POST":
		g.POST(path, h.Handle)
	}
}
