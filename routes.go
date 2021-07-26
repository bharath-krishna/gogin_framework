package main

import (
	_ "family-tree/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) Routes(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")
	r.GET("/oauth2-redirect.html", s.Oauth2RedirectUri)
	r.GET("/info", s.Info)
	url := ginSwagger.URL("http://localhost:8088/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	authUser := r.Group("/auth")
	{
		authUser.Use(s.AuthMiddleware())
		authUser.GET("/user", s.AuthWelcome)
	}
}
