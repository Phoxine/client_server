package main

import (
	_ "client_server/docs"
	"context"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample client server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.2.html

// @host localhost:1323
// @BasePath /api/v1

// @securitydefinitions.oauth2.password OAuth2Password
// @in header
// @name Authorization
// @tokenUrl /api/v1/auth/login

func main() {
	server, _ := InitializeServer()
	_ = server.Start(context.Background())
}
