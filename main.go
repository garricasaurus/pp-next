package main

import (
	"github.com/gin-gonic/gin"

	"okki.hu/garric/ppnext/consts"
	"okki.hu/garric/ppnext/controller"
)

func main() {
	// dummy change
	r := initRouter()
	scheduleBackgroundCleanup()
	r.Run(consts.Addr)
}

func initRouter() *gin.Engine {
	r := gin.Default()

	r.StaticFile("/favicon.ico", "./assets/favicon.ico")
	r.StaticFile("/scripts.js", "./assets/scripts.js")
	r.StaticFile("/materialize.min.css", "./assets/materialize.min.css")
	r.StaticFile("/materialize.min.css.map", "./assets/materialize.min.css.map")
	r.StaticFile("/materialize.min.js", "./assets/materialize.min.js")
	r.LoadHTMLGlob("templates/*")

	r.Use(controller.Auth())

	// public routes
	r.GET("/", controller.ShowLogin)
	r.GET("/login", controller.ShowLogin)
	r.POST("/login", controller.HandleLogin)
	r.GET("/logout", controller.HandleLogout)

	// protected routes
	prot := r.Group("/rooms", controller.Prot())
	prot.GET("/:room", controller.DisplayRoom)
	prot.GET("/:room/userlist", controller.UserList)
	prot.GET("/:room/results", controller.Results)
	prot.GET("/:room/events", controller.GetEvents)

	active := prot.Group("/", controller.Active())
	active.POST("/:room/vote", controller.AcceptVote)
	active.POST("/:room/reveal", controller.Reveal)
	active.POST("/:room/reset", controller.ResetRoom)

	return r
}
