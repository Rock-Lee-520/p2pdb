package event

//module name+topic name+event

//p2pdb-log EventType
const (
	p2pdbLogTopic = "p2pdb-log"
	LogTopic      = "log"
)

//p2pdb-store EventType
const (
	//DML
	StoreUpdateEvent    = "store_update_event"
	StoreSqlDeleteEvent = "store_delete_event"
	StoreSqlInsertEvent = "store_insert_event"
	//DDL table
	StoreCreateTableEvent   = "store_create_table_event"
	StoreAlterTableEvent    = "store_alter_table_event"
	StoreDropTableEvent     = "store_drop_table_event"
	StoreTruncateTableEvent = "store_truncate_table_event"
	//DDL database
	StoreDropDatabaseEvent   = "store_drop_database_event"
	StoreCreateDatabaseEvent = "store_create_database_event"
)
