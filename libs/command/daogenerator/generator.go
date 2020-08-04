package main

import (
	"html/template"
	"os"
	"strings"
)

type Generator struct {
	daoTmpl   string
	modelTmpl string
	output    string
	modelPkg  string
	processor *DataProcessor
	filter    []string
	genModel  bool
}

func (gen *Generator) createFile(context string, fileName string) {
	if !strings.HasSuffix(gen.output, "/") {
		gen.output = gen.output + "/"
	}

	if _, err := os.Stat(gen.output + context); os.IsNotExist(err) {
		os.MkdirAll(gen.output+context, 0744)
	}

	if _, err := os.Stat(gen.output + context + "/" + fileName); err == nil {
		os.Remove(gen.output + context + "/" + fileName)
	}
	_, err := os.Create(gen.output + context + "/" + fileName)
	if err != nil {
		panic(err)
	}
}

func (gen *Generator) generate() {
	var daotmpl *template.Template
	if gen.daoTmpl == "" {
		t, err := template.New("defaultDao").Funcs(TempFuncMap).Parse(defaultDaoTmpl)
		if err != nil {
			panic(err)
		}
		daotmpl = t
	} else {
		t, err := template.New("defaultDao").Funcs(TempFuncMap).ParseFiles(gen.daoTmpl)
		if err != nil {
			panic(err)
		}
		daotmpl = t
	}
	gen.GenDao(daotmpl)

	if gen.genModel == true {
		var modeltmpl *template.Template
		if gen.modelTmpl == "" {
			t, err := template.New("defaultModel").Parse(defaultModelTmpl)
			if err != nil {
				panic(err)
			}
			modeltmpl = t
		} else {
			t, err := template.New("defaultModel").ParseFiles(gen.modelTmpl)
			if err != nil {
				panic(err)
			}
			modeltmpl = t
		}
		gen.GenModel(modeltmpl)
	}
}

func (gen *Generator) GenDao(tmpl *template.Template) {
	for _, table := range gen.processor.Tables {
		filted := false
		for i, _ := range gen.filter {
			if table.TableName == strings.TrimSpace(gen.filter[i]) {
				filted = true
				break
			}
		}
		if filted {
			continue
		}
		fileName := table.TableName + "_dao.go"
		gen.createFile("dao", fileName)
		file, err := os.OpenFile(gen.output+"dao/"+fileName, os.O_RDWR, os.ModeAppend)
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := file.Close(); err != nil {
				panic(err)
			}
		}()

		ctx := gen.PrepareContext(table)
		err = tmpl.Execute(file, ctx)
		if err != nil {
			panic(err)
		}
	}
}

func (gen *Generator) GenModel(tmpl *template.Template) {
	for _, table := range gen.processor.Tables {
		filted := false
		for i, _ := range gen.filter {
			if table.TableName == strings.TrimSpace(gen.filter[i]) {
				filted = true
				break
			}
		}
		if filted {
			continue
		}
		fileName := table.TableName + ".go"
		gen.createFile("model", fileName)
		file, err := os.OpenFile(gen.output+"model/"+fileName, os.O_RDWR, os.ModeAppend)
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := file.Close(); err != nil {
				panic(err)
			}
		}()

		ctx := gen.PrepareContext(table)
		err = tmpl.Execute(file, ctx)
		if err != nil {
			panic(err)
		}
	}
}

func (gen *Generator) PrepareContext(table *Table) *Table {
	table.ModelPkg = gen.modelPkg
	return table
}
