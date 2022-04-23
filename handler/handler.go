package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, tpl, title string, data gin.H) {
	allData := gin.H{
		"title": title,
	}
	if data != nil && len(data) > 0 {
		for key, val := range data {
			allData[key] = val
		}
	}
	c.HTML(http.StatusOK, tpl, allData)
}
func renderNoData(c *gin.Context, tpl, title string) {
	render(c, tpl, title, nil)
}

func errHandler(c *gin.Context, err error) {
	c.String(http.StatusOK, err.Error())
	c.Abort()
}

func Index(c *gin.Context) {
	renderNoData(c, "index.html", "后台首页")
}
func Frontpage(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/m")
}
