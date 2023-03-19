package main

import (
	middleware "commerce-platform/middleware"
	"commerce-platform/modules"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./public", false)))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.Use(middleware.GetCoreResponseMiddleWare())
	r.Use(middleware.GetAuthenticationMiddleWare())

	modules.InitRouters(r)

	r.Run(":3000")
}
