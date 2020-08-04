package main

import (
	"fmt"
	"sort"

	"github.com/serenize/snaker"
)

type DataProcessor struct {
	Tables []*Table
}

func (p *DataProcessor) prepare() {
	db := conn()
	defer db.Close()

	showTables, err := db.Query("SHOW TABLES")
	defer showTables.Close()

	if err != nil {
		panic(err)
	}

	var tableNames []string
	for showTables.Next() {
		var tableName string
		showTables.Scan(&tableName)
		tableNames = append(tableNames, tableName)
	}

	for _, tableName := range tableNames {
		table := &Table{
			TableName: tableName,
			ModelName: snaker.SnakeToCamel(tableName),
		}
		p.Tables = append(p.Tables, table)

		rows, err := db.Query(fmt.Sprintf("select * from `%s`", tableName))
		defer rows.Close()

		if err != nil {
			panic(err)
		}

		types, err := rows.ColumnTypes()
		if err != nil {
			panic(err)
		}

		var cols []Column
		for _, tp := range types {

			nullable, _ := tp.Nullable()
			precision, scale, ok := tp.DecimalSize()
			var hasDecimal bool
			if ok {
				hasDecimal = true
			}
			length, _ := tp.Length()
			// value := tp.ScanType()

			column := Column{
				DatabaseTypeName: tp.DatabaseTypeName(),
				Name:             tp.Name(),
				PropName:         snaker.SnakeToCamel(tp.Name()),
				PropTypeName:     GetPropTypeName(tp.DatabaseTypeName()),
				Nullable:         nullable,
				HasDecimal:       hasDecimal,
				DecimalPrecision: precision,
				DecimalScale:     scale,
				Length:           length,
			}
			cols = append(cols, column)
		}
		sort.Sort(Columns(cols))
		table.Columns = cols
	}
}

func GetPropTypeName(databaseTypeName string) string {
	switch databaseTypeName {
	case fieldTypeDecimal:
		return scanTypeDecimal
	case fieldTypeTiny:
		return scanTypeBoolean
	case fieldTypeShort:
		return scanTypeInt
	case fieldTypeLong:
		return scanTypeInt
	case fieldTypeFloat:
		return scanTypeFloat32
	case fieldTypeDouble:
		return scanTypeFloat64
	case fieldTypeTimestamp:
		return scanTypeTime
	case fieldTypeLongLong:
		return scanTypeInt64
	case fieldTypeInt24:
		return scanTypeInt
	case fieldTypeDate:
		return scanTypeTime
	case fieldTypeTime:
		return scanTypeTime
	case fieldTypeDateTime:
		return scanTypeTime
	case fieldTypeYear:
		return scanTypeInt
	case fieldTypeVarCharBin:
		return scanTypebytes
	case fieldTypeBit:
		return scanTypeBoolean
	case fieldTypeBLOB:
		return scanTypeString
	case fieldTypeVarString:
		return scanTypeString
	case fieldTypeLongBLOB:
		return scanTypeString
	case fieldTypeTinyBLOBBin:
		return scanTypeString
	case fieldTypeTinyBLOB:
		return scanTypeString
	default:
		return scanTypeString
	}
}
