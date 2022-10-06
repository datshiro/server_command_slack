package slack_handler

import (
	"fmt"
	"server/internal/infras/server/consts"

	"github.com/gin-gonic/gin"
)

func NewRequest() Request {
	return &request{}
}

type Request interface {
	Bind(e *gin.Context) error
	Validate() error
	GetCommand() string
	GetChannelID() string
	GetText() string
}

type request struct {
	Command   string `form:"command"`
	ChannelID string `form:"channel_id"`
	Text      string `form:"text"`
}

func (r *request) GetCommand() string {
	return r.Command
}

func (r *request) GetChannelID() string {
	return r.ChannelID
}

func (r *request) GetText() string {
	return r.Text
}

func (r *request) Bind(e *gin.Context) error {
	return e.Bind(r)
}

func (r *request) Validate() error {
	if r.Command == "" {
		return consts.ErrorBadRequest.Detail(fmt.Errorf("Command cannot be empty"))
	}
	return nil
}
