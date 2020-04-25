package ws

import (
	"github.com/gin-gonic/gin"
)

type AccessTokenWs interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	FindById(context *gin.Context)
	UpdateExpirationTime(context *gin.Context)
}
