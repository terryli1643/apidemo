package main

import "html/template"

var tempFuncMap template.FuncMap

func init() {
	tempFuncMap = make(template.FuncMap)
	tempFuncMap["CheckField"] = CheckField
}

func CheckField(fields []Column, fieldName string) bool {
	for _, f := range fields {
		if f.Name == fieldName {
			return true
		}
	}
	return false
}
