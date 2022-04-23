package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, tpl string, data gin.H) {
	allData := gin.H{}
	for key, val := range data {
		allData[key] = val
	}
	c.HTML(http.StatusOK, tpl, allData)
}

func Index(c *gin.Context) {
	render(c, "index.html", gin.H{"title": "后台首页"})
}
func Frontpage(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/m")
}
