package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PublicController struct{}

// @Router /public/ping [get]
// @Description Ping
// @Tags Public
// @Success 200 {object} string
func (publicController *PublicController) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
