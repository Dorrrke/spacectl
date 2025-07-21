package shell

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

var currentDir string

func Start() {
	loadConfig()
	loadAliases()
	initCommands()

	var err error
	currentDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(getPromt()),
		prompt.OptionTitle("spacectl"),
	)
	p.Run()
}

func getPromt() string {
	return fmt.Sprintf("\033[1;32m%s\033[0m > ", currentDir)
}

func executor(input string) {
	input = strings.TrimSpace(input)
	if input == "" {
		return
	}

	args := strings.Split(input, " ")
	command := args[0]
	params := args[1:]

	runCommand(command, params)

	var err error
	currentDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}
