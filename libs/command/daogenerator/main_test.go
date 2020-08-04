package main

import (
	"testing"
)

func Test(t *testing.T) {
	DSN = "root:111111@/qpay?charset=utf8&parseTime=True&loc=Local"
	processor := &DataProcessor{}
	processor.prepare()

	generator := &Generator{
		processor: processor,
		output:    "./dao",
		modelPkg:  "./model",
	}
	generator.generate()
}
