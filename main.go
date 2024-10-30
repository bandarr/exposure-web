package main

import (
	"exposure-web/rfexposure"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func formHandler(c *gin.Context) {
	data := rfexposure.TestStub()
	c.HTML(http.StatusOK, "form.html", gin.H{
		"results": data,
	})
}

func submitHandler(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}
	// Handle form submission logic here
}

func main() {
	r := gin.Default()

	r.SetHTMLTemplate(template.Must(template.ParseFiles("form.html")))

	r.GET("/", formHandler)
	r.POST("/submit", submitHandler)
	r.Static("/static", "./static")

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
