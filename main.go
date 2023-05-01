package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "SELAM!!")
	})
	route.Run(":5555")
}

func loadTemplate() (*template.Template, error) {
	tmpl := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err := tmpl.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
