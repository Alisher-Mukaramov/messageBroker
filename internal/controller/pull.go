package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (c Controller) Pull(ctx *gin.Context) {
	var (
		httpCode int
		response interface{}
	)

	defer func() {
		if httpCode != http.StatusOK {
			ctx.AbortWithStatus(httpCode)
			return
		}
		ctx.Data(httpCode, "application/text; charset=utf-8", []byte(response.(string)))
	}()

	key := ctx.Param("key")
	sTimeout := ctx.Query("timeout")

	iTimeout := 0
	if sTimeout != "" {
		iTimeout, _ = strconv.Atoi(sTimeout)
	}

	if key == "" {
		httpCode = http.StatusNotFound
		return
	}

	ctxWithTimeout, _ := context.WithTimeout(ctx.Request.Context(), time.Second*time.Duration(iTimeout))

	response, ok := c.svc.Pull(key, ctxWithTimeout)
	if !ok {
		httpCode = http.StatusNotFound
		return
	}

	httpCode = http.StatusOK
}
