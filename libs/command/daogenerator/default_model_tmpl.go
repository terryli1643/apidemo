package main

var defaultModelTmpl = `package model
import (
	"time"

	"github.com/shopspring/decimal"
)

type {{.ModelName}} struct {
{{range $_, $col := .Columns }}
{{$col.PropName}}	{{$col.PropTypeName}} {{end}}
}

func (model {{.ModelName}}) TableName() string {
	return "{{.TableName}}"
}
`
