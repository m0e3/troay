package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GroupIndex(c *gin.Context) {
	render(c, "group.html", gin.H{"title": "分组列表"})
}
func GroupAdd(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		render(c, "group-add.html", gin.H{"title": "添加分组"})
		return
	}
}
func GroupEdit(c *gin.Context) {}
func GroupDel(c *gin.Context)  {}
func GroupSub(c *gin.Context)  {}
