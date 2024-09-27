package main

import (
	"database/sql"
	"fmt"
	"gen_mybatis/template/model"
	"gen_mybatis/utils"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

var (
	dsn             = "root:root@tcp(localhost:3306)/fission_activity"
	ProjectRootPath = "changle-fission-activity/activity/model"
)

func main() {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println("Error pinging the database:", err)
		return
	}
	// 查询所有表
	tableNames := selectAllTable(db)

	// 遍历表字段
	for _, tableName := range tableNames {
		rows, err := db.Query("SHOW COLUMNS FROM " + tableName)
		if err != nil {
			fmt.Println("Error fetching columns:", err)
			return
		}
		defer rows.Close()

		var columns []Column
		for rows.Next() {
			var col Column
			if err := rows.Scan(&col.Field, &col.ColType, &col.Null, &col.Key, &col.Default, &col.Extra); err != nil {
				fmt.Println("Error scanning row:", err)
				return
			}
			columns = append(columns, col)
		}

		// 组装struct数据
		modelStr, xmlModel := generateModel(tableName, columns)
		utils.WriteFile("genFiles/entity/"+tableName+"_entity.go", modelStr)

		// xml
		tableNameCamel := utils.ToCamelCase(tableName)
		//utils.WriteFile("mapper/"+tableNameCamel+"Mapper.xml", xmlStr)

		utils.WriteTemplateToFile("genFiles/mapper/"+tableNameCamel+"Mapper.xml", "template/xml_template.txt", xmlModel)

		// dao
		daoTemplateData := map[string]interface{}{
			"ProjectRootPath":     ProjectRootPath,
			"TableNameCamelLower": utils.ToLowerFirstChar(tableNameCamel),
			"TableNameCamelUpper": tableNameCamel,
		}
		utils.WriteTemplateToFile("genFiles/dao/"+tableName+"_mapper.go", "template/dao_template.txt", daoTemplateData)

		fmt.Println(tableName + "生成成功！")
	}

}

func selectAllTable(db *sql.DB) []string {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		fmt.Println("Error fetching columns:", err)
		return nil
	}
	defer rows.Close()

	var tableNames []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil
		}
		tableNames = append(tableNames, tableName)
	}
	return tableNames
}

func generateModel(tableName string, columns []Column) (string, model.XmlModel) {
	tableNameCamel := utils.ToCamelCase(tableName)

	var columnList []*model.ColumnInfo

	var entityBuilder strings.Builder
	entityBuilder.WriteString(fmt.Sprintf("package entity\n\nimport \"time\" \n\ntype %s struct {\n", capitalize(tableNameCamel+"Entity")))

	//var xmlResultMapBuilder strings.Builder
	//xmlResultMapBuilder.WriteString(fmt.Sprintf("    <resultMap id=\"BaseResultMap\"  tables=\"%s\">\n", tableName))
	//
	//var xmlSqlBuilder strings.Builder
	//xmlSqlBuilder.WriteString(fmt.Sprintf("    <sql id=\"Base_Column_List\">\n"))
	for _, column := range columns {
		fieldCamel := utils.ToCamelCase(column.Field)
		goType := utils.ConvertMySQLTypeToGoType(column.ColType)
		goXmlType := utils.ConvertMySQLTypeToGoTypeXml(column.ColType)

		entityBuilder.WriteString(fmt.Sprintf("    %s %s `json:\"%s\" gm:\"%s\"`\n", capitalize(fieldCamel), goType, column.Field, column.Field))

		//if index == 0 {
		//	xmlResultMapBuilder.WriteString(fmt.Sprintf("        <id column=\"%s\" langType=\"%s\"/>\n", column.Field, goXmlType))
		//	xmlSqlBuilder.WriteString(fmt.Sprintf("        %s", column.Field))
		//} else {
		//	xmlResultMapBuilder.WriteString(fmt.Sprintf("        <result column=\"%s\" langType=\"%s\"/>\n", column.Field, goXmlType))
		//	xmlSqlBuilder.WriteString(fmt.Sprintf(",%s", column.Field))
		//}

		columnInfo := &model.ColumnInfo{
			DbColumnField: column.Field,
			LangType:      goXmlType,
		}
		columnList = append(columnList, columnInfo)
	}
	//xmlResultMapBuilder.WriteString("    </resultMap>\n")
	//xmlSqlBuilder.WriteString(fmt.Sprintf("\n    </sql>\n"))

	//xmlFmt := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<!DOCTYPE mapper PUBLIC \"-//mybatis.org//DTD Mapper 3.0//EN\"\n        \"https://raw.githubusercontent.com/zhuxiujia/GoMybatis/master/mybatis-3-mapper.dtd\">\n\n<mapper>\n %s\n%s\n\n</mapper>"

	entityBuilder.WriteString("}\n")
	//return entityBuilder.String(), fmt.Sprintf(xmlFmt, xmlResultMapBuilder.String(), xmlSqlBuilder.String())
	return entityBuilder.String(), model.XmlModel{
		TableName:  tableName,
		ColumnList: columnList,
	}
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

// sql.NullString 用于处理可能为 NULL 的字符串
type Column struct {
	Field   string
	ColType string
	Null    string
	Key     string
	Default sql.NullString // 使用 sql.NullString
	Extra   string
}
