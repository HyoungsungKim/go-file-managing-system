package main

import (
	"fmt"

	"fileDB.com/src/internal/controller/APIv1"
	dbHandler "fileDB.com/src/internal/controller/DBHandler"
)

func main() {
	var (
		DB_SERVER_ADDRESS = "172.30.0.1"
		DB_SERVER_PORT    = "8090"
	)
	db := dbHandler.ConnectDB()
	router := APIv1.SetupRouter(db)

	router.LoadHTMLGlob("../testClient/*")
	router.MaxMultipartMemory = 8 << 20

	fString := fmt.Sprintf("%s:%s", DB_SERVER_ADDRESS, DB_SERVER_PORT)

	println("Listening on http://" + fString)
	router.Run(fString)
}
