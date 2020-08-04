package main

const (
	scanTypeFloat32 = "float32"
	scanTypeFloat64 = "float64"
	scanTypeInt     = "int"
	scanTypeInt64   = "int64"
	scanTypeString  = "string"
	scanTypebytes   = "bytes"
	scanTypeDecimal = "decimal.Decimal"
	scanTypeBoolean = "bool"
	scanTypeTime    = "time.Time"
)

type fieldType byte

const (
	fieldTypeDecimal     = "DECIMAL"
	fieldTypeTiny        = "TINYINT"
	fieldTypeShort       = "SMALLINT"
	fieldTypeLong        = "INT"
	fieldTypeFloat       = "FLOAT"
	fieldTypeDouble      = "DOUBLE"
	fieldTypeNULL        = "NULL"
	fieldTypeTimestamp   = "TIMESTAMP"
	fieldTypeLongLong    = "BIGINT"
	fieldTypeInt24       = "MEDIUMINT"
	fieldTypeDate        = "DATE"
	fieldTypeTime        = "TIME"
	fieldTypeDateTime    = "DATETIME"
	fieldTypeYear        = "YEAR"
	fieldTypeVarCharBin  = "VARBINARY"
	fieldTypeBit         = "BIT"
	fieldTypeBLOB        = "BLOB"
	fieldTypeVarString   = "VARCHAR"
	fieldTypeLongBLOB    = "LONGTEXT"
	fieldTypeBinary      = "BINARY"
	fieldTypeChar        = "CHAR"
	fieldTypeTinyBLOBBin = "TINYTEXT"
	fieldTypeTinyBLOB    = "TINYBLOB"
)
