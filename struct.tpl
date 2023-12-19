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
    {{.Name | snakecase}} := {{.Name}}{
      {{- range $index, $input := .Inputs}}  
      {{$input.Name | title}} :{{ if eq $input.Type "common.Address"}} common.HexToAddress(vLog.Topics[{{$index | add1 }}]) {{else if eq $input.Type "big.Int"}} vLog.Topics[{{$index | add1}}].(*bigInt), {{end}}
      {{- end}}
    }
  {{ end}}
}
