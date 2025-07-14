package helptext

var CustomHelpTemplate = `{{.Short}}

Доступные команды:
{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name 12}} {{.Short}}{{end}}{{end}}

Введите "help [команда]" для справки по конкретной команде.
`
