package shell

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Command struct {
	Handler     func(args []string)
	Description string
}

var customCommands = map[string]Command{}

func initCommands() {
	registerCommand("help", helpCommand, "Показать это сообщение")
	registerCommand("agent reboot", agentRebootCommand, "Перезагрузить агента")
	// Можно добавить еще комнады
}

func registerCommand(name string, handler func(args []string), description string) {
	customCommands[name] = Command{
		Handler:     handler,
		Description: description,
	}
}

func runCommand(cmd string, args []string) {
	full := cmd

	if len(args) > 0 {
		full = fmt.Sprintf("%s %s", cmd, strings.Join(args, " "))
	}

	if a, ok := aliases[cmd]; ok {
		full = a
		args = []string{}
		cmd = strings.Split(a, " ")[0]
	}

	switch cmd {
	case "cd":
		if len(args) == 0 {
			fmt.Println("Вы должны указать директорию")
			return
		}
		if err := os.Chdir(args[0]); err != nil {
			fmt.Println("Ошибка:", err.Error())
		}
		return
	case "export":
		if len(args) == 0 {
			fmt.Println("Пример: export VAR=value")
			return
		}
		for _, pair := range args {
			parts := strings.SplitN(pair, "=", 2)
			if len(parts) == 2 {
				os.Setenv(parts[0], parts[1])
			}
		}
		return
	case "alias":
		if len(args) < 2 {
			fmt.Println("Пример: alias name \"command string\"")
			return
		}
		aliases[args[0]] = strings.Join(args[1:], " ")
		saveAliases()
		return
	case "exit":
		os.Exit(0)
	}

	if cmdFunc, ok := customCommands[cmd]; ok {
		cmdFunc.Handler(args)
		return
	}

	if cmdFunc, ok := customCommands[full]; ok {
		cmdFunc.Handler([]string{})
		return
	}

	bin := exec.Command("bash", "-c", full)
	bin.Stdout = os.Stdout
	bin.Stderr = os.Stderr
	bin.Run()
}

func helpCommand(args []string) {
	fmt.Println("Доступные команды:")
	fmt.Println()

	for name, cmd := range customCommands {
		fmt.Printf(" %-20s - %s\n", name, cmd.Description)
	}
}

func agentRebootCommand(args []string) {
	fmt.Println("Перезагрузка агента...")
}
