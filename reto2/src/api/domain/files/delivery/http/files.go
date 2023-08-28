package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *FilesHTTPHandler) ReadAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := handler.usecase.ReadAll()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("error: %s", err.Error()))
		}
		ctx.JSON(http.StatusOK, res)
	}
}
