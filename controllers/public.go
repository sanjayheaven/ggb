package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PublicController struct{}

// @Summary Ping
// @Description Ping
// @Tags Public
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Router /public/ping [get]
func (publicController *PublicController) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
