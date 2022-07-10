package server

import (
	//"log"

	"time"

	"github.com/Rock-liyi/p2pdb-store/memory"
	"github.com/Rock-liyi/p2pdb-store/sql"
	"github.com/Rock-liyi/p2pdb-store/sqlite"
)

func CreateSqliteDatabase(DBName string, tableName string) *sqlite.Database {

	db := sqlite.NewDatabase(DBName)
	table := sqlite.NewTable(tableName, sql.NewPrimaryKeySchema(sql.Schema{
		{Name: "name2", Type: sql.Text, Nullable: false, Source: tableName},
		{Name: "email2", Type: sql.Text, Nullable: false, Source: tableName},
		{Name: "id", Type: sql.Int64, Nullable: false, Source: tableName},
		// {Name: "phone_numbers", Type: sql.JSON, Nullable: false, Source: tableName},
		// {Name: "created_at", Type: sql.Timestamp, Nullable: false, Source: tableName},
	}))
	// debug.Dump(table)
	db.AddTable(tableName, table)
	// ctx := sql.NewEmptyContext()
	// table.Insert(ctx, sql.NewRow("John Doe", "john@doe.com", 1))
	// table.Insert(ctx, sql.NewRow("John Doe", "john@doe.com", 2))
	// table.Insert(ctx, sql.NewRow("John Doe", "johnalt@doe.com", []string{}, time.Now()))
	// table.Insert(ctx, sql.NewRow("Jane Doe", "jane@doe.com", []string{}, time.Now()))
	// table.Insert(ctx, sql.NewRow("Evil Bob", "evilbob@gmail.com", []string{"555-666-555", "666-666-666"}, time.Now()))
	// db.DropTable(ctx, tableName)
	return db
}

func CreateMemoryDatabase(DBName string, tableName string) *memory.Database {

	db := memory.NewDatabase(DBName)
	table := memory.NewTable(tableName, sql.NewPrimaryKeySchema(sql.Schema{
		{Name: "name2", Type: sql.Text, Nullable: false, Source: tableName},
		{Name: "email2", Type: sql.Text, Nullable: false, Source: tableName},
		{Name: "id", Type: sql.Int64, Nullable: false, Source: tableName},
		{Name: "phone_numbers", Type: sql.JSON, Nullable: false, Source: tableName},
		{Name: "created_at", Type: sql.Timestamp, Nullable: false, Source: tableName},
	}))
	// debug.Dump(table)
	db.AddTable(tableName, table)
	ctx := sql.NewEmptyContext()
	table.Insert(ctx, sql.NewRow("John Doe", "john@doe.com", 1))
	table.Insert(ctx, sql.NewRow("John Doe", "john@doe.com", 2))
	table.Insert(ctx, sql.NewRow("John Doe", "johnalt@doe.com", []string{}, time.Now()))
	table.Insert(ctx, sql.NewRow("Jane Doe", "jane@doe.com", []string{}, time.Now()))
	table.Insert(ctx, sql.NewRow("Evil Bob", "evilbob@gmail.com", []string{"555-666-555", "666-666-666"}, time.Now()))
	return db
}
