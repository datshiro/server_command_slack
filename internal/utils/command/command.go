package command

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func ParsePlainCmdString(text string) []string {
	texts := strings.Split(text, "\n")[0] // eliminate last \n
	texts = strings.Trim(texts, " ")      // Trim both end whitespaces
	cmds := strings.Split(texts, " ")
	return cmds
}

func GetInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	log.Printf("text %v", text)
	return text, nil
}

// CmdExec Execute a command
func CmdExec(args ...string) (string, error) {

	baseCmd := args[0]
	cmdArgs := args[1:]

	log.Printf("Exec: %v", args)

	cmd := exec.Command(baseCmd, cmdArgs...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
