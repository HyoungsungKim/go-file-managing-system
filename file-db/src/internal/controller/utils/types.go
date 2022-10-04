package utils

type UserLogs struct {
	AccountId       string `json:"account_id"`
	LatestTimestamp string `json:"latest_timestamp"`
}

type UploadFormat struct {
	AccountId string
	FileName  string
	Signature string
	Type      string
	URI       string
	NFTtitle  string
	NFTId     string
	Copyright string
}

type RentalRequestFormat struct {
	AccountId    string `json:"account_id"`
	UserId       string `json:"user_id"`
	RequestorId  string `json:"requestor_id"`
	NFTId        string `json:"nft_id"`
	RentalPeriod string `json:"rental_period"`
	Timestamp    string `json:"timestamp"`
}
