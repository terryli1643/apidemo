package main

import (
	"flag"
	"log"
	"strings"
)

func main() {
	flag.StringVar(&DSN, "dsn", "", "data source name")
	output := flag.String("output", "", "output folder")
	modelPkg := flag.String("modelpkg", "", "model path")
	daoTmpl := flag.String("daotmpl", "", "template file path")
	modelTmpl := flag.String("modeltmpl", "", "template file path")
	filter := flag.String("filter", "", "filted tablen names, join by ','")
	genmodel := flag.Bool("genmodel", false, "generate model")
	tablePrefix := flag.String("tablePrefix", "", "table prefix")

	flag.Parse()

	log.Print("dsn:" + DSN)
	log.Print("daotmpl:" + *daoTmpl)
	log.Print("modeltmpl:" + *daoTmpl)
	log.Print("modelpkg:" + *modelPkg)
	log.Print("output:" + *output)
	log.Print("filter:" + *filter)
	log.Printf("genmodel:%t", *genmodel)
	log.Print("tablePrefix:" + *tablePrefix)

	processor := &DataProcessor{
		TablePrefix: *tablePrefix,
	}
	processor.prepare()

	generator := &Generator{
		daoTmpl:   *daoTmpl,
		modelTmpl: *modelTmpl,
		processor: processor,
		output:    *output,
		modelPkg:  *modelPkg,
		filter:    strings.Split(*filter, ","),
		genModel:  *genmodel,
	}

	generator.generate()
	log.Println("Finished")
}
