package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

var index_html = `
<html>
    <head>
        <title>JSONP DEMO</title>
    </head>
    <body>
        <script src="https://localhost:9999/JSONP?callback=alert&uid=1"></script>
    </body>
</html>
`

func main() {
	r := gin.Default()

	index := template.Must(template.New("index.html").Parse(index_html))
	r.SetHTMLTemplate(index)

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run(":10999")
}
