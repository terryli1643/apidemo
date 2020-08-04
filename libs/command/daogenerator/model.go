package main

type Table struct {
	TableName  string
	StructName string
	ModelPkg   string
	ModelName  string
	Columns    []Column
	Remark     string
}

type Columns []Column

type Column struct {
	DatabaseTypeName string
	Name             string
	PropName         string
	PropTypeName     string
	Nullable         bool
	HasDecimal       bool
	DecimalPrecision int64
	DecimalScale     int64
	Length           int64
}

func (c Columns) Len() int {
	return len(c)
}

func (c Columns) Less(i, j int) bool {
	return c[i].Name < c[j].Name
}

func (c Columns) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
