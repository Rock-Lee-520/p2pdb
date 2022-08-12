package entity

type DMLChangeType struct {
	INSERT string
	UPDATE string
	DELETE string
}

type DDLChangeType struct {
	CREATE          string
	ALTER           string
	DROP            string
	DDL_ACTION_TYPE DDL_ACTION_TYPE
}

type DDL_ACTION_TYPE struct {
	DATABASE string
	TABLE    string
}

type ChangeRowData interface {
	getDDLChangeType() *DDLChangeType
}
