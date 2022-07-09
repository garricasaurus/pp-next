package main

import (
	"github.com/gin-gonic/gin"
	"log"

	"okki.hu/garric/ppnext/controller"
)

func main() {
	r := initRouter()
	scheduleBackgroundCleanup()
	log.Fatal(r.Run())
}

func initRouter() *gin.Engine {
	r := gin.Default()

	r.StaticFile("/favicon.ico", "./assets/favicon.ico")
	r.StaticFile("/scripts.js", "./assets/scripts.js")
	r.StaticFile("/styles.css", "./assets/styles.css")
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

	api := r.Group("/rooms", controller.Api())
	api.GET("/:room/userlist", controller.UserList)
	api.GET("/:room/results", controller.Results)
	api.GET("/:room/events", controller.GetEvents)

	active := api.Group("/", controller.Active())
	active.POST("/:room/vote", controller.AcceptVote)
	active.POST("/:room/reveal", controller.Reveal)
	active.POST("/:room/reset", controller.ResetRoom)

	_ = r.SetTrustedProxies(nil)

	return r
}
