package shell

import (
	promt "github.com/c-bata/go-prompt"
)

func completer(d promt.Document) []promt.Suggest {
	word := d.GetWordBeforeCursor()
	suggestions := []promt.Suggest{
		{Text: "cd", Description: "Перейти в директорию"},
		{Text: "export", Description: "Установить переменную окружения"},
		{Text: "exit", Description: "Выйти из shell"},
		{Text: "help", Description: "Помощь"},
		{Text: "alias", Description: "Добавить псеводним"},
	}

	for name, cmd := range customCommands {
		suggestions = append(suggestions, promt.Suggest{
			Text:        name,
			Description: cmd.Description,
		})
	}
	for name := range aliases {
		suggestions = append(suggestions, promt.Suggest{
			Text:        name,
			Description: "Псевдоним",
		})
	}

	return promt.FilterHasPrefix(suggestions, word, true)
}
