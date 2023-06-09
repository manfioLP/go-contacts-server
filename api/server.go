package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BasicPath(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Server is healthy")
}

func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
