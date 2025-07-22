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
	SubCommands map[string]*Command
}

var customCommands = map[string]*Command{}

func initCommands() {
	registerCommand("agent reboot", agentRebootCommand, "Перезагрузить агента")
	registerCommand("agent stop", agentRebootCommand, "Остановка агента")
	registerCommand("glint-launcher stop", agentRebootCommand, "Остановка лаунчера")
	registerCommand("glint-launcher reboot", agentRebootCommand, "Перезагрузка лаунчера")
	// Можно добавить еще комнады
}

func registerCommand(path string, handler func(args []string), description string) {
	parts := strings.Split(path, " ")
	curr := customCommands

	for i, part := range parts {
		if i == len(parts)-1 {
			curr[part] = &Command{
				Handler:     handler,
				Description: description,
				SubCommands: map[string]*Command{},
			}
		} else {
			if _, ok := curr[part]; !ok {
				curr[part] = &Command{
					Description: "",
					SubCommands: map[string]*Command{},
				}
			}
			curr = curr[part].SubCommands
		}
	}
}

func runCommand(input string, args []string) {
	full := input
	if len(args) > 0 {
		full = fmt.Sprintf("%s %s", input, strings.Join(args, " "))
	}

	if a, ok := aliases[input]; ok {
		full = a
		args = []string{}
		input = strings.Split(a, " ")[0]
	}

	if input == "help" {
		fmt.Println("Доступные команды:")
		for name, cmd := range customCommands {
			if cmd.Description != "" {
				fmt.Printf("  %-15s - %s\n", name, cmd.Description)
			} else {
				fmt.Printf("  %-15s\n", name)
			}

			// если есть подкоманды — выведем и их
			if cmd.SubCommands != nil {
				for subName, sub := range cmd.SubCommands {
					fmt.Printf("    %-13s - %s\n", name+" "+subName, sub.Description)
				}
			}
		}
		return
	}

	if cmdEntry, ok := customCommands[input]; ok {
		if cmdEntry.SubCommands != nil {
			// Подкоманда передана?
			if len(args) == 0 || args[0] == "-h" || args[0] == "--help" || args[0] == "-help" {
				fmt.Printf("Подкоманды для '%s':\n", input)
				for name, sub := range cmdEntry.SubCommands {
					fmt.Printf("  %-15s - %s\n", name, sub.Description)
				}
				return
			}

			subCmd := args[0]
			subArgs := args[1:]

			if subEntry, ok := cmdEntry.SubCommands[subCmd]; ok {
				subEntry.Handler(subArgs)
				return
			}

			fmt.Printf("Неизвестная подкоманда: %s\n", subCmd)
			return
		}
		// Обычная команда без подкоманд
		cmdEntry.Handler(args)
		return
	}

	switch input {
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

	bin := exec.Command("bash", "-c", full)
	bin.Stdout = os.Stdout
	bin.Stderr = os.Stderr
	bin.Run()
}

// func helpCommand(args []string) {
// 	fmt.Println("Доступные команды:")
// 	fmt.Println()
// 	printHelp(customCommands, "")
// }

// func printHelp(cmds map[string]*Command, prefix string) {
// 	for name, cmd := range cmds {
// 		full := strings.TrimSpace(prefix + " " + name)
// 		if cmd.Description != "" {
// 			fmt.Printf(" %-25s - %s\n", full, cmd.Description)
// 		}
// 		if len(cmd.SubCommands) > 0 {
// 			printHelp(cmd.SubCommands, full)
// 		}
// 	}
// }

func agentRebootCommand(args []string) {
	fmt.Println("Перезагрузка агента...")
}
