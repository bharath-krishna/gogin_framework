package main

import (
	_ "family-tree/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) Routes(r *gin.Engine) {
	r.GET("/info", s.Info)
	r.LoadHTMLGlob("templates/*")
	r.GET("/oauth2-redirect.html", s.Oauth2RedirectUri)
	r.GET("/logout", s.Logout)

	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	protected := r.Group("/").Use(s.AuthMiddleware())
	{
		protected.GET("/", s.Welcome)
		protected.GET("/user", s.AuthWelcome)
	}
}
