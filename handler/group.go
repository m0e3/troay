package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m0e3/troay/defs"
)

func GroupIndex(c *gin.Context) {
	renderNoData(c, "group.html", "分组列表")
}
func GroupAdd(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		renderNoData(c, "group-add.html", "添加分组")
		return
	}
	var form defs.GroupAdd
	if err := c.ShouldBind(&form); err != nil {
		errHandler(c, err)
		return
	}

}
func GroupEdit(c *gin.Context) {}
func GroupDel(c *gin.Context)  {}
func GroupSub(c *gin.Context)  {}
