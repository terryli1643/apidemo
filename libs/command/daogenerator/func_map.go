package main

import "html/template"

var TempFuncMap template.FuncMap

func init() {
	TempFuncMap = make(template.FuncMap)
	TempFuncMap["CheckField"] = CheckField
}

func CheckField(fields []Column, fieldName string) bool {
	for _, f := range fields {
		if f.Name == fieldName {
			return true
		}
	}
	return false
}
