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
	Notch   string `json:"notch"`
}

//go:embed rotors.json
var data []byte

func main() {
	f, err := os.Create("../../rotors_gen.go")
	if err != nil {
		log.Fatalf("create: %s", err.Error())
	}
	defer f.Close()

	var rotors []JSONConfig
	if err := json.Unmarshal(data, &rotors); err != nil {
		log.Fatalf("unmarshal: %s", err.Error())
	}

	funcs := template.FuncMap{
		"keyMap": keyMap,
	}
	t := template.Must(template.New("config").Funcs(funcs).Parse(tpl))
	if err := t.Execute(f, rotors); err != nil {
		log.Fatalf("execute: %s", err.Error())
	}
}

var errTooFewKeys = func(mapping string) error { return fmt.Errorf("too few keys in mapping: %s", mapping) }

func keyMap(keys string) (map[enigma.Key]enigma.Key, error) {
	alphas := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if len(keys) != len(alphas) {
		return nil, errTooFewKeys(keys)
	}

	out := make(map[enigma.Key]enigma.Key, len(keys))
	for i := 0; i < len(alphas); i++ {
		out[enigma.Key(alphas[i])] = enigma.Key(keys[i])
	}
	return out, nil
}

const tpl = `package enigma

var (
	{{- range .}}
	Rotor{{.Name}} = rotorConfig{
		mapping: map[Key]Key{
			{{- range $k, $v := keyMap .Mapping}}
			{{$k}}: {{$v}},
			{{- end}}
		},
		notch: {{.Notch}},
	}
	{{- end }}
)
`
