switch event.Name {
  {{range .}}
  case "{{.Name}}":
  {{end}}
}

