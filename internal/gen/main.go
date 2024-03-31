package main

import (
	"bytes"
	"html/template"
	"log"
	"os"
)

func main() {
	buf := bytes.Buffer{}

	t := template.Must(template.New("bindings").Parse(bindingsTmpl))
	err := t.Execute(&buf, struct {
		Main    []Variant
		MapVars []MapVariant
	}{
		Main:    getVariants(),
		MapVars: getMapVariants(),
	})
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("bindings_gen.go", buf.Bytes(), 0660)
	if err != nil {
		log.Fatal(err)
	}

}

var bindingsTmpl = `package chh

import (
    "time"

	"github.com/ClickHouse/ch-go/proto"
)
{{range .Main}}
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
// -- type Map[K]V

{{range .MapVars}}
func (r *ColResults) BindMap{{.Key}}To{{.Value}}(name string, valPtr interface{}) {
	col := proto.NewMap[{{.GoKey}}, {{.GoValue}}](&proto.Col{{.Key}}{}, &proto.Col{{.Value}}{})
	r.Bind(name, col, &BindingMap[{{.GoKey}}, {{.GoValue}}]{col: col, valPtr: valPtr})
}
{{end}}
`
