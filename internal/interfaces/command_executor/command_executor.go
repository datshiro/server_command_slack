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
	fields := make([]*slack.TextBlockObject, 0)
	blocks := []slack.Block{}

	for index, line := range strings.Split(out, "\n") {
		if line == "" {
			continue
		}
		fields = append(fields, slack.NewTextBlockObject("mrkdwn", line, false, false))
		if index%9 == 0 { // For every 10 line wrap into 1 section
			fieldsSection := slack.NewSectionBlock(nil, fields, nil)
			blocks = append(blocks, fieldsSection)
			fields = make([]*slack.TextBlockObject, 0)
		}
	}
	fieldsSection := slack.NewSectionBlock(nil, fields, nil)
	blocks = append(blocks, fieldsSection)

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
