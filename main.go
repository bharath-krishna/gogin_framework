package main

import (
	"os"
)

// @title Go Gin API framework - Boilerplate code
// @version 3.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
// @securitydefinitions.oauth2.implicit
// @authorizationurl http://localhost:8099/auth/realms/demo/protocol/openid-connect/auth
func main() {
	app := newApp()
	app.Run(os.Args)
}
