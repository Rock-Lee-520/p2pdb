package event

//event+module name+topic name

//p2pdb-log EventType
const (
	p2pdbLogTopic = "p2pdb-log"
	LogTopic      = "log"
)

//p2pdb-store EventType
const (
	//DML
	StoreSqlUpdateEvent = "store_sql_update_event"
	StoreSqlDeleteEvent = "store_sql_delete_event"
	StoreSqlInsertEvent = "store_sql_insert_event"
	//DDL table
	StoreSqlCreateTableEvent   = "store_sql_create_table_event"
	StoreSqlAlterTableEvent    = "store_sql_alter_table_event"
	StoreSqlDropTableEvent     = "store_sql_drop_table_event"
	StoreSqlTruncateTableEvent = "store_sql_truncate_table_event"
	//DDL database
	StoreSqlDropDatabaseEvent   = "store_sql_drop_database_event"
	StoreSqlCreateDatabaseEvent = "store_sql_create_database_event"
)
