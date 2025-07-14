package shell

import (
	"fmt"
	"os"
	"strings"

	"github.com/Dorrrke/spacectl/cmd"
	"github.com/chzyer/readline"
)

func Start() {
	completer := readline.NewPrefixCompleter(
		readline.PcItem("reboot", readline.PcItem("agent"), readline.PcItem("glint")),
	)

	rl, err := readline.NewEx(&readline.Config{
		Prompt:       ">>",
		HistoryFile:  "/tmp/readline.tmp",
		AutoComplete: completer,
	})

	if err != nil {
		panic(err)
	}

	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		args := strings.Fields(line)
		cmd.RootCmd.SetArgs(args)
		if err := cmd.RootCmd.Execute(); err != nil {
			fmt.Fprint(os.Stderr, "Ошибка: ", err, "\n")
		}
	}
}
