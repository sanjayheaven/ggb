package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PublicController struct{}

func (publicController *PublicController) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
