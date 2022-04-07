package main

//go:generate go run .

import (
	_ "embed"
	"encoding/json"
	"html/template"
	"log"
	"os"

	"github.com/vmorsell/enigma/enigma/generators"
)

type JSONConfig struct {
	Name    string `json:"name"`
	Mapping string `json:"mapping"`
}

//go:embed reflectors.json
var data []byte

func main() {
	f, err := os.Create("../../reflectors_gen.go")
	if err != nil {
		log.Fatalf("create: %v", err)
	}
	defer f.Close()

	var reflectors []JSONConfig
	if err := json.Unmarshal(data, &reflectors); err != nil {
		log.Fatalf("unmarshal: %v", err)
	}

	funcs := template.FuncMap{
		"charMap": generators.CharMap,
	}
	t := template.Must(template.New("config").Funcs(funcs).Parse(tpl))
	if err := t.Execute(f, reflectors); err != nil {
		log.Fatalf("execute: %v", err)
	}
}

const tpl = `package enigma

var (
	{{- range .}}
	Reflector{{.Name}} = ReflectorType{
		mapping: map[Char]Char{
			{{- range $k, $v := charMap .Mapping}}
			{{$k}}: {{$v}},
			{{- end}}
		},
	}
	{{- end}}

	AllReflectorTypes = []ReflectorType{
		{{- range .}}
		Reflector{{.Name}},
		{{- end}}
	}
)
`
