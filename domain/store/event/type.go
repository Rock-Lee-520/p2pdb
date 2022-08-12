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
	StoreUpdateEvent = "store_update_event"
	StoreDeleteEvent = "store_delete_event"
	StoreInsertEvent = "store_insert_event"
	//DDL table
	StoreCreateTableEvent   = "store_create_table_event"
	StoreAlterTableEvent    = "store_alter_table_event"
	StoreDropTableEvent     = "store_drop_table_event"
	StoreTruncateTableEvent = "store_truncate_table_event"
	//DDL database
	StoreDropDatabaseEvent   = "store_drop_database_event"
	StoreCreateDatabaseEvent = "store_create_database_event"
)

var StoreEventType = [...]string{
	//DML
	StoreUpdateEvent,
	StoreDeleteEvent,
	StoreInsertEvent,
	//DDL table
	StoreCreateTableEvent,
	StoreAlterTableEvent,
	StoreDropTableEvent,
	StoreTruncateTableEvent,
	StoreDropDatabaseEvent,
	StoreCreateDatabaseEvent,
}

//p2pdb-pubsub EventType
const (
	PubsubStartEvent = "pubsub_start_event"
)
