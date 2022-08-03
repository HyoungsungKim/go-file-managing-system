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
	collectionGroup(db, v1.Group("/collection"))

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
			"account_id": data["account_id"],
			"file_name":  data["file_name"],
			"signature":  data["signature"],
			"type":       data["type"],
			"URI":        data["URI"],
			"size":       data["size"],
		})
		fmt.Println(data)
		doc, _ := json.Marshal(data)
		fmt.Println(string(doc))

		uploadData := utils.UploadFormat{
			AccountID: data["account_id"].(string),
			FileName:  data["file_name"].(string),
			Signature: data["signature"].(string),
			Type:      data["type"].(string),
			URI:       data["URI"].(string),
			Size:      data["size"].(float64),
		}

		DBHandler.Upload(db, uploadData)

		c.String(http.StatusOK, string(doc))
	})
}

func collectionGroup(db *sql.DB, router *gin.RouterGroup) {
	type ImageList struct {
		AccountID string   `json:"account_id"`
		Signature string   `json:"signature"`
		FileName  []string `json:"file_name"`
		URI       []string `json:"URI"`
		Type      []string `json:"type"`
	}

	router.GET("/:accountID", func(c *gin.Context) {
		var imageList ImageList
		var signature string
		accountID := c.Param("accountID")
		imageList.AccountID = accountID

		rows := DBHandler.ViewImages(db, accountID)
		defer rows.Close()

		for rows.Next() {
			var uploadedFormat utils.UploadFormat
			if err := rows.Scan(
				&uploadedFormat.AccountID,
				&uploadedFormat.FileName,
				&uploadedFormat.Signature,
				&uploadedFormat.Type,
				&uploadedFormat.URI,
				&uploadedFormat.Size,
			); err != nil {
				panic(err)
			}
			signature = uploadedFormat.Signature
			imageList.FileName = append(imageList.FileName, uploadedFormat.FileName)
			imageList.URI = append(imageList.URI, uploadedFormat.URI)
			imageList.Type = append(imageList.Type, uploadedFormat.Type)
		}
		imageList.Signature = signature
		/*
			jsonImageList, err := json.Marshal(imageList)
			if err != nil {
				panic(err)
			}
		*/
		//log.Print(imageList)
		c.JSON(http.StatusOK, imageList)
	})
}
