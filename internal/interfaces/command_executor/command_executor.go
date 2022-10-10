package command_executor

import (
	"context"
	"fmt"
	commandUtils "server/internal/utils/command"
	"strings"

	"github.com/slack-go/slack"
)

func New() Executor {
	return &executor{}
}

type Executor interface {
	Run(command string, ctx context.Context) (string, error)
}

type executor struct {
}

func (e *executor) Run(command string, ctx context.Context) (string, error) {
	cmds := commandUtils.ParsePlainCmdString(command)
	out, err := commandUtils.CmdExec(cmds...)
	if err != nil {
		return "", err
	}
	fmt.Printf("Command output is %s\n", out)
	return out, nil
}

func ParseSlackOutput(out string) slack.Message {
	blocks := []slack.Block{}
	for _, line := range strings.Split(out, "\n") {
		if line == "" {
			continue
		}
		field := slack.NewTextBlockObject("mrkdwn", line, false, false)
		fieldsSection := slack.NewSectionBlock(field, nil, nil)
		blocks = append(blocks, fieldsSection)
	}
	return slack.NewBlockMessage(blocks...)
}

// func buildSlackBlocks(msg string) SlackBlock {
// 	return SlackBlock{
// 		Type: "section",
// 		Text: SlackText{
// 			Type: "mrkdwn",
// 			Text: msg,
// 		},
// 	}
// }
