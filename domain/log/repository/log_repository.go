package repository

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	PO "github.com/Rock-liyi/p2pdb/domain/log/repository/po"
	orm "github.com/Rock-liyi/p2pdb/infrastructure/core/orm"
)

var DB *orm.CreateDBFactory = new(orm.CreateDBFactory)

func init() {
	//settings  use  the default  information database
	DB.SetIsInternalStore(true)
}

func CreateNodeTable() bool {
	db := DB.InitDB()
	db.Migrator().CreateTable(&PO.Node{})
	return true
}

func CreateDataTable(tableName string) bool {
	db := DB.InitDB()
	db.Migrator().CreateTable(&PO.Data{})
	return true
}

func CreateLinkTable(tableName string) bool {
	db := DB.InitDB()
	db.Migrator().CreateTable(&PO.Data{})
	return true
}

type Table struct {
	TableName    string
	SelectFields string
	FieldNames   []string
	FieldTypes   []string
	FieldDefault []string
	FieldIsNull  []bool
	DataType     reflect.Type
}

func NewTable(db *sql.DB, object interface{}, name string) Table {
	// get fields
	rows, err := db.Query("SELECT COLUMN_NAME, data_type,column_default,is_nullable FROM information_schema.COLUMNS WHERE TABLE_NAME = '" + name + "' ORDER BY ordinal_position;")
	if err != nil {
		panic(err)
	}
	tb := Table{TableName: name}
	var field, typ, coldf, isnil string
	for rows.Next() {
		rows.Scan(&field, &typ, &coldf, &isnil)
		tb.FieldNames = append(tb.FieldNames, field)
		tb.FieldTypes = append(tb.FieldTypes, tabletypes[typ])
		tb.FieldDefault = append(tb.FieldDefault, coldf)
		tb.FieldIsNull = append(tb.FieldIsNull, isnil == "YES")
	}
	tb.SelectFields = `"` + strings.Join(tb.FieldNames, `","`) + `"`

	// get struct type
	iValue := reflect.ValueOf(object)
	for iValue.Kind() == reflect.Ptr || iValue.Kind() == reflect.Interface {
		iValue = iValue.Elem()
	}
	if iValue.Kind() == reflect.Struct || iValue.Kind() == reflect.Map {
		tb.DataType = iValue.Type()
	} else {
		var sf []reflect.StructField = make([]reflect.StructField, len(tb.FieldNames))
		for i := 0; i < len(tb.FieldNames); i++ {
			name := strings.ToLower(tb.FieldNames[i])
			sf[i] = reflect.StructField{
				Name: strings.ToTitle(name),
				Type: tableReflectTypes[tb.FieldTypes[i]],
				Tag:  reflect.StructTag(fmt.Sprintf("alias:\"%s\" json:\"%s\"", name, name)),
			}
		}
		tb.DataType = reflect.StructOf(sf)
	}
	return tb
}

var tabletypes = map[string]string{
	"boolean":                     "Bool",
	"integer":                     "Int",
	"real":                        "Float",
	"character":                   "Byte",
	"character varying":           "String",
	"text":                        "String",
	"date":                        "Date",
	"time without time zone":      "Time",
	"timestamp without time zone": "DateTime",
}

var tableReflectTypes = map[string]reflect.Type{
	"Bool":     reflect.TypeOf((*bool)(nil)).Elem(),
	"Int":      reflect.TypeOf((*int)(nil)).Elem(),
	"Float":    reflect.TypeOf((*float64)(nil)).Elem(),
	"Byte":     reflect.TypeOf((*byte)(nil)).Elem(),
	"String":   reflect.TypeOf((*string)(nil)).Elem(),
	"Date":     reflect.TypeOf((*time.Time)(nil)).Elem(),
	"Time":     reflect.TypeOf((*time.Time)(nil)).Elem(),
	"DateTime": reflect.TypeOf((*time.Time)(nil)).Elem(),
}
