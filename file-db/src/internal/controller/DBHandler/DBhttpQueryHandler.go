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
		INSERT INTO "metadata" ("account_id", "file_name", "signature","type", "uri", "nft_title", "nft_id", "copyright") VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := db.Exec(uploadStmt,
		uploadFormat.AccountID,
		uploadFormat.FileName,
		uploadFormat.Signature,
		uploadFormat.Type,
		uploadFormat.URI,
		uploadFormat.NFTtitle,
		uploadFormat.NFTID,
		uploadFormat.Copyright,
	)
	checkError(err, uploadStmt)
}

func SelectAllImages(db *sql.DB, accountID string) *sql.Rows {
	selectImagesStmt := `
		SELECT * FROM metadata WHERE account_id = $1
	`

	rows, err := db.Query(selectImagesStmt, accountID)
	checkError(err, selectImagesStmt)

	return rows
}

func SelectImageByNFTID(db *sql.DB, NFTID string) *sql.Rows {
	selectImageByNFTIDstmt := `
		SELECT * FROM metadata WHERE nft_id = $1
	`

	rows, err := db.Query(selectImageByNFTIDstmt, NFTID)
	checkError(err, selectImageByNFTIDstmt)

	return rows
}
