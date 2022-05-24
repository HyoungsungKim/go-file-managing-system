package main

import (
	"database/sql"
	"testing"

	dbHandler "fileDB.com/src/internal/controller/DBHandler"
)

func TestConnectDB(t *testing.T) {
	db := dbHandler.ConnectDB()
	t.Logf("Connected to database")

	createStmt := `create table students(
		name varchar (20) Primary key,
		roll integer	
	)`

	_, err := db.Exec(createStmt)
	checkError(err, db, createStmt, t)

	selectStmt := `select * from students`
	_, err = db.Exec(selectStmt)
	checkError(err, db, selectStmt, t)

	insertStmt := `insert into "students"("name", "roll") values('John', 1)`
	_, err = db.Exec(insertStmt)
	checkError(err, db, insertStmt, t)

	insertDynStmt := `insert into "students"("name", "roll") values($1, $2)`
	_, err = db.Exec(insertDynStmt, "Jane", 2)
	checkError(err, db, insertDynStmt, t)

	_, err = db.Exec(selectStmt)
	checkError(err, db, selectStmt, t)

	db.Exec(`drop table students`)

	db.Close()
}

func checkError(err error, db *sql.DB, stmt string, t *testing.T) {
	if err != nil {
		//panic(err)
		db.Exec(`drop table students`)
		t.Log(err)
		t.Fatal(stmt)
	}
}
