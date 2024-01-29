package main

import (
	"embed"

	"github.com/sanjayheaven/ggb/cmd"
)

//go:embed cmd/*
//go:embed configs/*.go configs/config.example.yaml
//go:embed deployments/*
//go:embed githooks/*
//go:embed internal/*
//go:embed scripts/*
//go:embed tools/*
//go:embed web/template/*
//go:embed .air.toml .gitignore go.mod go.sum main.go Makefile
var EmbedFs embed.FS

// @title Go Gin Boilerplate API
// @description This is a swagger api document for Go Gin Boilerplate.
// @version v1.0.0

// @contact.name Dorvan
// @contact.url https://github.com/sanjayheaven

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and JWT token.
func main() {
	cmd.Execute(EmbedFs)
}
