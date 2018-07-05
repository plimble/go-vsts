package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

const html = `
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="%s/%s/%s/%s git https://%s.visualstudio.com/%s/_git/%s">
</head>
<body>
Nothing to see here!
</body>
</html>
`

func render(c *gin.Context) {
	result := fmt.Sprintf(html,
		c.Request.Host, c.Param("account"), c.Param("project"), c.Param("repo"),
		c.Param("account"), c.Param("project"), c.Param("repo"),
	)
	c.Data(200, "text/html", []byte(result))
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.GET("/:account/:project/:repo/*path", func(c *gin.Context) {
		render(c)
	})

	r.GET("/:account/:project/:repo", func(c *gin.Context) {
		render(c)
	})

	log.Fatal(r.Run(":8000"))
}
