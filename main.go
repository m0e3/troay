package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/m0e3/troay/handler"
)

func main() {
	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery())
	app.Static("/static", "./static")
	app.LoadHTMLGlob("./templates/*.html")
	app.GET("/", handler.Frontpage)
	g := app.Group("/m")
	{
		g.GET("/", handler.Index)
	}
	gg := g.Group("/group")
	{
		gg.GET("/", handler.GroupIndex)
		gg.GET("/add", handler.GroupAdd)
		gg.POST("/add", handler.GroupAdd)
		gg.GET("/edit/:id", handler.GroupEdit)
		gg.POST("/edit/:id", handler.GroupEdit)
		gg.GET("/del/:id", handler.GroupDel)
		gg.GET("/sub/:id", handler.GroupSub)
	}
	log.Fatal(app.Run(":59527"))
}
