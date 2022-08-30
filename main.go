package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func upload(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	files := form.File["files"]

	for _, file := range files {
		filename := filepath.Join("./uploadfile", filepath.Base(file.Filename))
		log.Println(file.Filename)

		if err := ctx.SaveUploadedFile(file, filename); err != nil {
			ctx.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}
	}
	ctx.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	router.StaticFile("/", "./public")
	router.Static("/file", "./uploadfile")
	router.POST("/upload", upload)

	router.Run(":8000")
}
