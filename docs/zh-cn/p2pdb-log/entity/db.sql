-- node definition

CREATE TABLE "node" (
	node_id TEXT NOT NULL,
	node_type TEXT,
	lamport_clock INTEGER,
	receiving_timestamp INTEGER,
	receiving_date TEXT,
	sending_date TEXT,
	sending_timestamp INTEGER,
	last_node_id TEXT,
	created_at TEXT,
	updated_at TEXT,
	deleted_at TEXT
);

-- link definition

CREATE TABLE "link" (
	link_id TEXT,
	last_node_id TEXT,
	node_id TEXT,
	created_at TEXT,
	updated_at TEXT,
	deleted_at TEXT
);



-- "object" definition

CREATE TABLE "object" (
	object_id TEXT NOT NULL,
	node_id TEXT NOT NULL,
	operation TEXT,
	propertie TEXT,
	context TEXT,
	created_at TEXT,
	updated_at TEXT,
	deleted_at TEXT
);