package controller

import (
	"bucket/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetFile(c *gin.Context) {

	bucket_name, _ := c.Params.Get("bucket")
	filepathname, _ := c.Params.Get("filename")

	key := filepath.Join(bucket_name, filepathname)

	// bucket_name, _ := c.Params.Get("bucket-name")
	// key = bucket_name + key
	filename := filepath.Base(key)

	data := service.Download(key)

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Header("Content-Type", http.DetectContentType(data))
	c.Header("Content-Length", fmt.Sprintf("%d", len(data)))
	c.Writer.Write(data) //the memory take up 1.2~1.7G
}

func Upload(c *gin.Context) {
	// sess, _ := session.NewSession()

	bucket_name, isName := c.Params.Get("bucket")
	if !isName {
		bucket_name = "abc"
	}

	// uploader := s3manager.NewUploader(sess)
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Println(err)
	}

	path, err := service.Upload(file, header, bucket_name)
	// filename := header.Filename

	file_name := filepath.Base(path)

	hostname := os.Getenv("HOST")

	downloadPath := hostname + "/" + "api/buckets" + "/" + bucket_name + "/" + file_name

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to upload file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"filepath": downloadPath,
	})
}
