package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strconv"

	"github.com/yigitoo/random_profile_picture/cmd"

	"github.com/gin-gonic/gin"
)

//go:embed assets/* templates/*
var f embed.FS

const PORT uint16 = 5555

func main() {
	router := gin.Default()
	templ := template.Must(template.New("").ParseFS(f, "templates/*.tmpl"))
	router.SetHTMLTemplate(templ)

	// example: /public/assets/images/example.png
	router.StaticFS("/public", http.FS(f))

	router.GET("/", func(ctx *gin.Context) {
		location := url.URL{Path: "/png"}
		ctx.Redirect(301, location.RequestURI())
	})
	router.GET("/png", func(ctx *gin.Context) {
		file_name := cmd.NewProfile("png")
		ctx.JSON(http.StatusOK, gin.H{
			"profile_link":      "/public/assets/" + file_name,
			"profile_full_link": "http://localhost:" + strconv.Itoa(int(PORT)) + "/public/assets/" + file_name,
		})
	})

	router.GET("favicon.ico", func(ctx *gin.Context) {
		file, _ := f.ReadFile("assets/favicon.ico")
		ctx.Data(
			http.StatusOK,
			"image/x-icon",
			file,
		)
	})

	router.GET("/hi", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello bro!!")
	})
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hi bro!!")
	})

	if PORT > 0 {
		router.Run(":" + strconv.Itoa(int(PORT)))
	}

	fmt.Println("Program finished.")
}
