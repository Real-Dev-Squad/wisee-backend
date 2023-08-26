package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	r := gin.Default()
	r.SetHTMLTemplate(html)

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})

	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	r.Run(":8080")
}
