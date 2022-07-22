package main

import (
	"fmt"
	"os"

	"fileServer.com/FileServer/src/internal/controller"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	var (
		SERVER_ADDRESS = os.Getenv("SERVER_ADDRESS")
		SERVER_PORT    = os.Getenv("SERVER_PORT")

		// For integer port
		//SERVER_PORT, _  = strconv.Atoi(os.Getenv("SERVER_PORT"))
	)

	router := controller.SetupRouter()
	//router.LoadHTMLGlob("../templates/*")
	router.MaxMultipartMemory = 8 << 20
	fString := fmt.Sprintf("%s:%s", SERVER_ADDRESS, SERVER_PORT)
	println("Listening on http://" + fString)
	router.Run(fString)
}

func loadEnv() {
	err := godotenv.Load("/app/.env")
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
