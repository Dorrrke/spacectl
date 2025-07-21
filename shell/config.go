package shell

import (
	"encoding/json"
	"fmt"
	"os"
)

var aliases = map[string]string{}

func loadAliases() {
	data, err := os.ReadFile("config/aliases.json")
	if err != nil {
		return
	}
	json.Unmarshal(data, &aliases)
}

func saveAliases() {
	data, _ := json.MarshalIndent(aliases, "", "  ")
	os.WriteFile("config/aliases.json", data, 0644)
}

func loadConfig() {
	if _, err := os.Stat("config/aliases.json"); err != nil {
		err := os.MkdirAll("config", 0755)
		if err != nil {
			fmt.Println("Ошибка создания config:", err.Error())
		}

		_ = os.WriteFile("config/aliases.json", []byte("{}"), 0644)
		_ = os.WriteFile("config/config.json", []byte("{}"), 0644)
	}
}
