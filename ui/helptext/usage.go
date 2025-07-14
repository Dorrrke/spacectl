package helptext

var CustomUsageTemplate = `
В интерактивном режиме просто введите нужную команду:

Доступные команды:
{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name 12}} {{.Short}}{{end}}{{end}}
`
