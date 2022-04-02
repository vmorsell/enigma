package main

//go:generate go run .

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/vmorsell/enigma/enigma"
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
		"keyMap": keyMap,
	}
	t := template.Must(template.New("config").Funcs(funcs).Parse(tpl))
	if err := t.Execute(f, reflectors); err != nil {
		log.Fatalf("execute: %v", err)
	}
}

var errTooFewKeys = func(mapping string) error { return fmt.Errorf("too few keys in mapping: %s", mapping) }

func keyMap(keys string) (map[enigma.Key]enigma.Key, error) {
	alphas := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if len(keys) != len(alphas) {
		return nil, errTooFewKeys(keys)
	}

	out := make(map[enigma.Key]enigma.Key, len(alphas))
	for i := 0; i < len(alphas); i++ {
		out[enigma.Key(alphas[i])] = enigma.Key(keys[i])
	}
	return out, nil
}

const tpl = `package enigma

var (
	{{- range .}}
	Reflector{{.Name}} = reflectorConfig{
		mapping: map[Key]Key{
			{{- range $k, $v := keyMap .Mapping}}
			{{$k}}: {{$v}},
			{{- end}}
		},
	}
	{{- end}}
)
`
