package main

import "MyCar/internal/app"

// @title MyCar API
// @version 1.0
// @description API for MyCar application
// @host localhost:8080
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	app.Run()
}
