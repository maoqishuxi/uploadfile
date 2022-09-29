package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
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
			ctx.String(http.StatusBadRequest, "uploadfile file err: %s", err.Error())
			return
		}
	}
	// ctx.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	ctx.Redirect(http.StatusMovedPermanently, "/file")
}

func getfilelist(ctx *gin.Context) {
	files, err := os.ReadDir("./uploadfile")
	if err != nil {
		ctx.String(http.StatusBadRequest, "not exist file: %s", err.Error())
	}

	filename := make([]string, 0)
	for _, file := range files {

		filename = append(filename, filepath.Join("/file", file.Name()))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": filename,
	})
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.MaxMultipartMemory = 8 << 20
	//router.StaticFile("/", "./public")
	router.Use(static.Serve("/", static.LocalFile("./public", false)))
	router.StaticFS("/file", http.Dir("./uploadfile"))
	router.POST("/uploadfile", upload)
	router.GET("/filelist", getfilelist)

	//router.Run(":8000")
	router.RunTLS(":8000", "./julai/julai.fun.pem", "./julai/julai.fun.key")
}
