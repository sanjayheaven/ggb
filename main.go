package main

import "github.com/sanjayheaven/ggb/cmd"

// @title Go Gin Boilerplate API
// @description This is a swagger api document for Go Gin Boilerplate.
// @version v1.0.0

// @contact.name Dorvan
// @contact.url https://github.com/sanjayheaven

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.

func main() {
	cmd.Execute()
}
