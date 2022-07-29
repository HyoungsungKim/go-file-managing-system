package DBHandler

import (
	"database/sql"

	"fileDB.com/src/internal/controller/utils"
)

func resetDB(db *sql.DB) {
	db.Exec(`drop table metadata`)
}

func Upload(db *sql.DB, uploadFormat utils.UploadFormat) {
	uploadStmt := `
		INSERT INTO "metadata" ("account_id", "file_name", "signature","type", "uri", "size") VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := db.Exec(uploadStmt,
		uploadFormat.AccountID,
		uploadFormat.FileName,
		uploadFormat.Signature,
		uploadFormat.Type,
		uploadFormat.URI,
		uploadFormat.Size,
	)
	checkError(err, uploadStmt)
}

func ViewImages(db *sql.DB, accountID string) *sql.Rows {
	viewImageStmt := `
		SELECT * FROM metadata WHERE account_id = $1
	`

	rows, err := db.Query(viewImageStmt, accountID)
	checkError(err, viewImageStmt)

	return rows
}
