package impl

import (
	"github.com/gin-gonic/gin"
	. "github.com/guidomantilla/bookstore_oauth-api/core/service"
)

type DefaultAccessTokenWs struct {
	accessTokenService AccessTokenService
}

func NewDefaultAccessTokenWs(accessTokenService AccessTokenService) *DefaultAccessTokenWs {
	return &DefaultAccessTokenWs{
		accessTokenService: accessTokenService,
	}
}

func (accessTokenWs *DefaultAccessTokenWs) Create(context *gin.Context) {

}

func (accessTokenWs *DefaultAccessTokenWs) Update(context *gin.Context) {}

func (accessTokenWs *DefaultAccessTokenWs) Delete(context *gin.Context) {}

func (accessTokenWs *DefaultAccessTokenWs) FindById(context *gin.Context) {}

func (accessTokenWs *DefaultAccessTokenWs) Find(context *gin.Context) {}

func (accessTokenWs *DefaultAccessTokenWs) UpdateExpirationTime(context *gin.Context) {}
