package DBHandler

import (
	"database/sql"
	"fmt"
)

func generateTables(db *sql.DB) {
	createaMetadataTable(db)
}

func createaMetadataTable(db *sql.DB) {
	createStmt := `create table IF NOT EXISTS "metadata" (
		userId 		text Primary key,
		title 		text,
		description text
	)`

	_, err := db.Exec(createStmt)
	checkError(err, createStmt)
}

func checkError(err error, stmt string) {
	if err != nil {
		fmt.Print(stmt)
		panic(err)
	}
}

/*
func checkTableExist(db *sql.DB, tableName string) bool {
	_, tableCheck := db.Exec("select * from " + tableName)
	if tableCheck != nil {
		return true
	} else {
		panic(tableCheck)
	}
}
*/