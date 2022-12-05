package DBHandler

import (
	"database/sql"
	"fmt"
)

func resetDB(db *sql.DB) {
	db.Exec(`drop table metadata`)
	db.Exec(`drop table rental_request`)
	db.Exec(`drop table user_logs`)
}

func generateTables(db *sql.DB) {
	createMetadataTable(db)
	createRentalRequestTable(db)
	createUserLogsTable(db)
}

func createUserLogsTable(db *sql.DB) {
	createStmt := `create table IF NOT EXISTS "user_logs" (
		account_id 			text,
		latest_timestamp	text,
		PRIMARY KEY (account_id)
	)`

	_, err := db.Exec(createStmt)
	checkError(err, createStmt)
}

func createMetadataTable(db *sql.DB) {
	createStmt := `create table IF NOT EXISTS "metadata" (
		owner_id	text,
		account_id 	text,
		file_name 	text,
		signature	text,
		type 		text,
		URI    		text,
		nft_title 	text,
		nft_id		text,
		copyright 	text,
		uci 		text
	)`

	_, err := db.Exec(createStmt)
	checkError(err, createStmt)
}

func createRentalRequestTable(db *sql.DB) {
	createStmt := `create table IF NOT EXISTS "rental_request" (
		account_id 		text,
		user_id 		text,
		nft_id			text,
		requestor_id    text,
		rental_period	text,
		timestamp	 	text
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
