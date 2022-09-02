package APIv1

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"fileDB.com/src/internal/controller/DBHandler"
	"fileDB.com/src/internal/controller/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()
	router.NoRoute(ReverseProxy)

	router.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{
				"http://localhost:3010",
				"http://172.30.0.1:8090",
				"http://172.30.0.1:3000",
			},
			AllowMethods: []string{"GET", "POST"},
			MaxAge:       12 * time.Hour,
		},
	))

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

func ReverseProxy(c *gin.Context) {
	//https://story.tomasen.org/a-better-practice-mixing-gin-with-next-js-and-nextauth-on-the-side-4f9d1fb9e486

	remote, _ := url.Parse("http://localhost:3000")

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL = c.Request.URL
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

func uploadGroup(db *sql.DB, router *gin.RouterGroup) {
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
