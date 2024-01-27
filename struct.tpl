{{ range . }}
type {{ .Name }} struct {
 {{- range .Inputs}} 
  {{.Name | title}} {{.Type}}  `json:"{{.Name}}"` 
  {{- end}}
}
{{ end }}

switch event.Name {
  {{- range .}}
  case "{{- .Name}}":
    {{ if .IsFetchLogData}}
    decodeLog, err := contractABI.Unpack(event.Name, vLog.Data)
    if err != nil {
      log.Fatal("unable to decode log")
    }
    {{end}}
    {{.Name | snakecase}} := {{.Name}}{
      {{- range .Inputs}}
        {{.Name | title}} : {{.InitValue}},
      {{- end}}
    }
  

    // do stuff

    return
  {{end}}
}


