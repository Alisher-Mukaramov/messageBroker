package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c Controller) Append(ctx *gin.Context) {
	var (
		httpCode int
	)

	defer func() {
		ctx.AbortWithStatus(httpCode)
	}()

	key := ctx.Param("key")
	value := ctx.Query("v")

	if key == "" || value == "" {
		httpCode = http.StatusBadRequest
		return
	}

	c.svc.Append(key, value)
	httpCode = http.StatusOK
}
