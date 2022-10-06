package slack_handler

import (
	"net/http"
	"server/internal/interfaces"
	"server/internal/interfaces/command_executor"

	"github.com/gin-gonic/gin"
)

func NewHandler() interfaces.Handler {
	return &handler{
		NewRequest: NewRequest,
		method:     http.MethodPost,
		path:       "slack",
	}
}

type handler struct {
	method     string
	path       string
	NewRequest func() Request
}

func (h *handler) Handle(ctx *gin.Context) {
	request := h.NewRequest()
	if err := request.Bind(ctx); err != nil {
		return
	}

	if err := request.Validate(); err != nil {
		return
	}

	exe := command_executor.New()
	out, err := exe.Run(request.GetText(), ctx.Request.Context())
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	msg := command_executor.ParseSlackOutput(out)
	ctx.JSON(200, msg)
}

func (h *handler) GetMethod() string {
	return h.method
}

func (h *handler) GetPath() string {
	return h.path
}
