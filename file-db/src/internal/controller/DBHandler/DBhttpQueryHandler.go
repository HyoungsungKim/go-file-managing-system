package DBHandler

import (
	"database/sql"

	"fileDB.com/src/internal/controller/utils"
)

func resetDB(db *sql.DB) {
	db.Exec(`drop table metadata`)
}

func Upload(db *sql.DB, uploadFormant utils.UploadFormat) {
	uploadStmt := `
		insert into "metadata" ("userid", "title", "description") values ($1, $2, $3)
	`

	_, err := db.Exec(uploadStmt, uploadFormant.UserId, uploadFormant.Title, uploadFormant.Description)
	checkError(err, uploadStmt)
}
