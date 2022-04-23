package handler

import "github.com/gin-gonic/gin"

func GroupIndex(c *gin.Context) {
	render(c, "group.html", gin.H{"title": "分组列表"})
}
func GroupAdd(c *gin.Context)  {}
func GroupEdit(c *gin.Context) {}
func GroupDel(c *gin.Context)  {}
func GroupSub(c *gin.Context)  {}
