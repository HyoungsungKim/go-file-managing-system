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
			AllowMethods: []string{"GET", "POST", "PUT"},
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
	indexGroup(db, v1.Group("/")) // 되나...?
	uploadGroup(db, v1.Group("/upload"))
	collectionGroup(db, v1.Group("/collection"))
	requestGroup(db, v1.Group("/request"))

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

func indexGroup(db *sql.DB, router *gin.RouterGroup) {
	router.GET("/user-logs/:accountId", func(c *gin.Context) {
		var userLogs utils.UserLogs

		accountId := c.Param("accountId")
		rows := DBHandler.SelectUserLogs(db, accountId)
		defer rows.Close()

		for rows.Next() {
			if err := rows.Scan(
				&userLogs.AccountId,
				&userLogs.LatestTimestamp,
			); err != nil {
				panic(err)
			}
		}
		c.JSON(http.StatusOK, userLogs)
	})

	router.PUT("/user-logs/:accountId", func(c *gin.Context) {

		body := c.Request.Body
		value, err := ioutil.ReadAll(body)
		if err != nil {
			panic(err)
		}

		//accountId := c.Param("accountId")
		var data map[string]interface{}
		json.Unmarshal([]byte(value), &data)
		c.JSON(http.StatusOK, gin.H{
			"account_id":       data["account_id"],
			"latest_timestamp": data["latest_timestamp"],
		})
		fmt.Println(data)
		doc, _ := json.Marshal(data)
		fmt.Println(string(doc))

		userLogs := utils.UserLogs{
			AccountId:       data["account_id"].(string),
			LatestTimestamp: data["latest_timestamp"].(string),
		}

		DBHandler.PutUserLogs(db, userLogs)

		c.String(http.StatusOK, string(doc))
	})

	router.GET("/rental-logs/:accountId", func(c *gin.Context) {
		type RentalLogsList struct {
			AccountId     string   `json:"account_id"`
			UserIds       []string `json:"user_ids"`
			NFTIds        []string `json:"nft_ids"`
			RequestorIds  []string `json:"requestor_ids"`
			RentalPeriods []string `json:"rental_periods"`
			Timestamps    []string `json:"timestamps"`
		}

		var rentalLogsList RentalLogsList

		rentalLogsList.AccountId = c.Param("accountId")
		rows := DBHandler.SelectRentalRequestByAccountId(db, rentalLogsList.AccountId)

		defer rows.Close()
		for rows.Next() {
			var rentalLogsFormat utils.RentalRequestFormat
			if err := rows.Scan(
				&rentalLogsFormat.AccountId,
				&rentalLogsFormat.UserId,
				&rentalLogsFormat.NFTId,
				&rentalLogsFormat.RequestorId,
				&rentalLogsFormat.RentalPeriod,
				&rentalLogsFormat.Timestamp,
			); err != nil {
				panic(err)
			}
			rentalLogsList.UserIds = append(rentalLogsList.UserIds, rentalLogsFormat.UserId)
			rentalLogsList.NFTIds = append(rentalLogsList.NFTIds, rentalLogsFormat.NFTId)
			rentalLogsList.RequestorIds = append(rentalLogsList.RequestorIds, rentalLogsFormat.RequestorId)
			rentalLogsList.RentalPeriods = append(rentalLogsList.RentalPeriods, rentalLogsFormat.RentalPeriod)
			rentalLogsList.Timestamps = append(rentalLogsList.Timestamps, rentalLogsFormat.Timestamp)
		}
		/*
			jsonImageList, err := json.Marshal(imageList)
			if err != nil {
				panic(err)
			}
		*/
		//log.Print(imageList)
		c.JSON(http.StatusOK, rentalLogsList)
	})
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
			"NFTtitle":   data["NFTtitle"],
			"NFT_id":     data["NFT_id"],
			"Copyright":  data["Copyright"],
		})
		fmt.Println(data)
		doc, _ := json.Marshal(data)
		fmt.Println(string(doc))

		uploadData := utils.UploadFormat{
			AccountId: data["account_id"].(string),
			FileName:  data["file_name"].(string),
			Signature: data["signature"].(string),
			Type:      data["type"].(string),
			URI:       data["URI"].(string),
			NFTtitle:  data["NFTtitle"].(string),
			NFTId:     data["NFT_id"].(string),
			Copyright: data["Copyright"].(string),
		}

		DBHandler.InsertMetadata(db, uploadData)

		c.String(http.StatusOK, string(doc))
	})
}

func collectionGroup(db *sql.DB, router *gin.RouterGroup) {
	type ImageMetadataList struct {
		AccountId  string   `json:"account_id"`
		Signature  string   `json:"signature"`
		FileNames  []string `json:"file_names"`
		URIs       []string `json:"URIs"`
		Types      []string `json:"types"`
		NFTtitles  []string `json:"NFTtitles"`
		NFTIds     []string `json:"NFT_ids"`
		Copyrights []string `json:"copyrights"`
	}

	router.GET("/:accountId", func(c *gin.Context) {
		var imageList ImageMetadataList
		var signature string
		accountId := c.Param("accountId")
		imageList.AccountId = accountId

		rows := DBHandler.SelectAllImages(db, accountId)
		defer rows.Close()

		for rows.Next() {
			var uploadedFormat utils.UploadFormat
			if err := rows.Scan(
				&uploadedFormat.AccountId,
				&uploadedFormat.FileName,
				&uploadedFormat.Signature,
				&uploadedFormat.Type,
				&uploadedFormat.URI,
				&uploadedFormat.NFTtitle,
				&uploadedFormat.NFTId,
				&uploadedFormat.Copyright,
			); err != nil {
				panic(err)
			}
			signature = uploadedFormat.Signature
			imageList.FileNames = append(imageList.FileNames, uploadedFormat.FileName)
			imageList.URIs = append(imageList.URIs, uploadedFormat.URI)
			imageList.Types = append(imageList.Types, uploadedFormat.Type)
			imageList.NFTtitles = append(imageList.NFTtitles, uploadedFormat.NFTtitle)
			imageList.NFTIds = append(imageList.NFTIds, uploadedFormat.NFTId)
			imageList.Copyrights = append(imageList.Copyrights, uploadedFormat.Copyright)
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

func requestGroup(db *sql.DB, router *gin.RouterGroup) {
	type ImageMetadata struct {
		AccountId string `json:"account_id"`
		Signature string `json:"signature"`
		FileName  string `json:"file_name"`
		URI       string `json:"URI"`
		Type      string `json:"type"`
		NFTtitle  string `json:"NFTtitle"`
		NFTId     string `json:"NFT_id"`
		Copyright string `json:"copyright"`
	}

	router.GET("/:nftId", func(c *gin.Context) {
		var imageMetadata ImageMetadata
		var signature string
		nftId := c.Param("nftId")
		imageMetadata.NFTId = nftId

		rows := DBHandler.SelectImageByNFTId(db, nftId)
		defer rows.Close()

		for rows.Next() {
			var uploadedFormat utils.UploadFormat
			if err := rows.Scan(
				&uploadedFormat.AccountId,
				&uploadedFormat.FileName,
				&uploadedFormat.Signature,
				&uploadedFormat.Type,
				&uploadedFormat.URI,
				&uploadedFormat.NFTtitle,
				&uploadedFormat.NFTId,
				&uploadedFormat.Copyright,
			); err != nil {
				panic(err)
			}

			signature = uploadedFormat.Signature

			imageMetadata.AccountId = uploadedFormat.AccountId
			imageMetadata.FileName = uploadedFormat.FileName
			imageMetadata.URI = uploadedFormat.URI
			imageMetadata.Type = uploadedFormat.Type
			imageMetadata.NFTtitle = uploadedFormat.NFTtitle
			imageMetadata.NFTId = uploadedFormat.NFTId
			imageMetadata.Copyright = uploadedFormat.Copyright
			imageMetadata.Signature = signature

			if imageMetadata.Copyright == "unlockable content" {
				imageMetadata.URI = "404Images/lockedContent.png"
				imageMetadata.FileName = "locked file"
				imageMetadata.AccountId = "locked"
			}
		}

		/*
			jsonImageList, err := json.Marshal(imageList)
			if err != nil {
				panic(err)
			}
		*/
		//log.Print(imageList)
		c.JSON(http.StatusOK, imageMetadata)
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
			"account_id":    data["account_id"],
			"user_id":       data["user_id"],
			"requestor_id":  data["requestor_id"],
			"nft_id":        data["NFT_id"],
			"rental_period": data["rental_period"],
			"timestamp":     data["timestamp"],
		})

		fmt.Println(data)
		doc, _ := json.Marshal(data)
		fmt.Println(string(doc))

		uploadData := utils.RentalRequestFormat{
			AccountId:    data["account_id"].(string),
			UserId:       data["user_id"].(string),
			NFTId:        data["NFT_id"].(string),
			RequestorId:  data["requestor_id"].(string),
			RentalPeriod: data["rental_period"].(string),
			Timestamp:    data["timestamp"].(string),
		}

		DBHandler.InsertRentalRequest(db, uploadData)

		c.String(http.StatusOK, string(doc))
	})
}
