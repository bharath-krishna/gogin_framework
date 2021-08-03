package main

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// Welcome godoc
// @Summary
// @Description
// @Tags Welcome
// @Accept  json
// @Produce  json
// @Produce  application/json
// @Success 200 {object} map[string]string
// @Success 400 {object} map[string]string
// @Router / [get]
func (s *Server) Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"X-Forwarded-Access-Token": c.GetHeader("X-Forwarded-Access-Token"),
	})
	return
}

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
// @Router /user [get]
// @Security
func (s *Server) AuthWelcome(c *gin.Context) {
	user := c.MustGet("user")
	c.JSON(http.StatusOK, user)
}

func (s *Server) Oauth2RedirectUri(c *gin.Context) {
	// c.Request.Response.Header.Add("content-type", "text/html")
	c.HTML(http.StatusOK, "oauth2-redirect.html", nil)
}

func (s *Server) Logout(c *gin.Context) {
	q := url.Values{}
	q.Set("rd", "http://localhost:8099/auth/realms/demo/protocol/openid-connect/logout?redirect_uri=http://localhost:8888")
	location := url.URL{Path: "/oauth2/sign_out", RawQuery: q.Encode()}
	s.logger.Debug(location.RequestURI())
	c.Redirect(http.StatusTemporaryRedirect, location.RequestURI())
	return
}
