package model

type XmlModel struct {
	TableName  string
	ColumnList []*ColumnInfo
}

type ColumnInfo struct {
	DbColumnField string
	LangType      string
}
