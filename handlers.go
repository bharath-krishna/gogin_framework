package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Info godoc
// @Summary info page
// @Description This is the info page
// @Tags Info
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /info [get]
func (s *Server) Info(c *gin.Context) {
	info := map[string]string{"version": version}
	c.JSON(http.StatusOK, info)
}

// AuthWelcome godoc
// @Summary auth handler
// @Description auth handler
// @Tags User
// @Produce json
// @Produce application/json
// @Success 200 {object} User
// @Success 400 {object} map[string]string
// @Router /auth/user [get]
// @Security
func (s *Server) AuthWelcome(c *gin.Context) {
	user := c.MustGet("user")
	c.JSON(http.StatusOK, user)
}

func (s *Server) Oauth2RedirectUri(c *gin.Context) {
	// c.Request.Response.Header.Add("content-type", "text/html")
	c.HTML(http.StatusOK, "oauth2-redirect.html", nil)
}
