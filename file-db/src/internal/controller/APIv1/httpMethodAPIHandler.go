package APIv1

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"fileDB.com/src/internal/controller/DBHandler"
	"fileDB.com/src/internal/controller/utils"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Db server pong",
		})
	})

	router.GET("/test", func(c *gin.Context) {
		title := "main test page"
		c.HTML(http.StatusOK, "index.html", gin.H{
			"page": title,
		})
	})

	v1 := router.Group("/")
	uploadGroup(db, v1.Group("/upload"))

	return router
}

func uploadGroup(db *sql.DB, router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		title := "upload test page"
		c.HTML(http.StatusOK, "upload.html", gin.H{
			"page": title,
		})
	})

	router.POST("/submit", func(c *gin.Context) {
		body := c.Request.Body
		value, err := ioutil.ReadAll(body)
		if err != nil {
			panic(err)
		}

		var data map[string]interface{}
		json.Unmarshal([]byte(value), &data)
		c.JSON(http.StatusOK, gin.H{
			"userId":      data["userId"],
			"title":       data["age"],
			"description": data["description"],
		})
		fmt.Println(data)
		doc, _ := json.Marshal(data)
		fmt.Println(string(doc))

		uploadData := utils.UploadFormat{UserId: data["userId"].(string), Title: data["title"].(string), Description: data["description"].(string)}

		DBHandler.Upload(db, uploadData)

		c.String(http.StatusOK, string(doc))
	})
}
