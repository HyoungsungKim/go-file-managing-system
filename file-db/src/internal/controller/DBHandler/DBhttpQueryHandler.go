package DBHandler

import (
	"database/sql"

	"fileDB.com/src/internal/controller/utils"
)

func SelectUserLogs(db *sql.DB, accountId string) *sql.Rows {
	selectUserLogsStmt := `
		SELECT * FROM user_logs WHERE account_id = $1
	`

	rows, err := db.Query(selectUserLogsStmt, accountId)
	checkError(err, selectUserLogsStmt)

	return rows
}

func PutUserLogs(db *sql.DB, userLogs utils.UserLogs) {
	putUserLogsStmt := `
		INSERT INTO user_logs ("account_id", "latest_timestamp") VALUES ($1, $2)
		ON CONFLICT (account_id) DO UPDATE SET latest_timestamp = $2
	`

	_, err := db.Exec(putUserLogsStmt, userLogs.AccountId, userLogs.LatestTimestamp)
	checkError(err, putUserLogsStmt)

}

func InsertMetadata(db *sql.DB, uploadFormat utils.UploadFormat) {
	insertMetadataStmt := `
		INSERT INTO "metadata" ("account_id", "file_name", "signature","type", "uri", "nft_title", "nft_id", "copyright") VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := db.Exec(insertMetadataStmt,
		uploadFormat.AccountId,
		uploadFormat.FileName,
		uploadFormat.Signature,
		uploadFormat.Type,
		uploadFormat.URI,
		uploadFormat.NFTtitle,
		uploadFormat.NFTId,
		uploadFormat.Copyright,
	)
	checkError(err, insertMetadataStmt)
}

func SelectAllImages(db *sql.DB, accountId string) *sql.Rows {
	selectImagesStmt := `
		SELECT * FROM metadata WHERE account_id = $1
	`

	rows, err := db.Query(selectImagesStmt, accountId)
	checkError(err, selectImagesStmt)

	return rows
}

func SelectImageByNFTId(db *sql.DB, NFTId string) *sql.Rows {
	selectImageByNFTIdStmt := `
		SELECT * FROM metadata WHERE nft_id = $1
	`

	rows, err := db.Query(selectImageByNFTIdStmt, NFTId)
	checkError(err, selectImageByNFTIdStmt)

	return rows
}

func InsertRentalRequest(db *sql.DB, rentalRequestFormat utils.RentalRequestFormat) {
	uploadStmt := `
	INSERT INTO "rental_request" ("account_id", "user_id", "requestor_id", "nft_id", "rental_period", "timestamp") VALUES ($1, $2, $3, $4, $5, $6)
`

	_, err := db.Exec(uploadStmt,
		rentalRequestFormat.AccountId,
		rentalRequestFormat.UserId,
		rentalRequestFormat.NFTId,
		rentalRequestFormat.RequestorId,
		rentalRequestFormat.RentalPeriod,
		rentalRequestFormat.Timestamp,
	)
	checkError(err, uploadStmt)
}

func SelectRentalRequestByAccountId(db *sql.DB, accountId string) *sql.Rows {
	selectRentalRequestByAccountIdStmt := `
	SELECT * FROM rental_request WHERE account_id = $1
`

	rows, err := db.Query(selectRentalRequestByAccountIdStmt, accountId)
	checkError(err, selectRentalRequestByAccountIdStmt)

	return rows
}

func SelectRentalRequestByNFTId(db *sql.DB, NFTId string) *sql.Rows {
	selectRentalRequestByNFTIdstmt := `
	SELECT * FROM rental_request WHERE nft_id = $1
`

	rows, err := db.Query(selectRentalRequestByNFTIdstmt, NFTId)
	checkError(err, selectRentalRequestByNFTIdstmt)

	return rows
}
