package main

import (
	"bytes"
	"html/template"
	"log"
	"os"
)

func main() {
	buf := bytes.Buffer{}

	data := getVariants()
	t := template.Must(template.New("bindings").Parse(bindgingsTmpl))
	err := t.Execute(&buf, data)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("bindings_gen.go", buf.Bytes(), 0660)
	if err != nil {
		log.Fatal(err)
	}
}

var bindgingsTmpl = `
package chh

import (
    "time"

	"github.com/ClickHouse/ch-go/proto"
)
{{range .}}
// -- type {{ .Name }}

func (r *ColResults) Bind{{.Name}}(name string, valPtr interface{}) {
	col := &proto.{{.Type}}{}
	r.Bind(name, col, &binding{{.Type}}{col: col, valPtr: valPtr})
}

type binding{{.Type}} struct {
	col    *proto.{{.Type}}
	valPtr interface{}
}

func (b *binding{{.Type}}) Append() {
	s, _ := b.valPtr.(*[]{{.GoName}})
	appendToSlice(s, b.col)
}

func (b *binding{{.Type}}) Write() {
	if b.col.Rows() > 0 {
		s, _ := b.valPtr.(*{{.GoName}})
		*s = b.col.Row(0)
	}
}
{{end}}
`
