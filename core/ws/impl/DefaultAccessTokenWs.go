package impl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/guidomantilla/bookstore_common-lib/common/exception"
	. "github.com/guidomantilla/bookstore_oauth-api/core/model"
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

	var accessToken AccessToken
	if err := context.ShouldBindJSON(&accessToken); err != nil {
		exception := BadRequestException("error unmarshalling request json to object", err)
		context.JSON(exception.Code, exception)
		return
	}

	if exception := accessTokenWs.accessTokenService.Create(&accessToken); exception != nil {
		context.JSON(exception.Code, exception)
		return
	}

	//isPublic := context.GetHeader("X-Public") == "true"
	//marshaledUser := MarshallUser(&user, isPublic)

	context.JSON(http.StatusCreated, accessToken)
}

func (accessTokenWs *DefaultAccessTokenWs) Update(context *gin.Context) {

	id := context.Param("id")

	var accessToken AccessToken
	if err := context.ShouldBindJSON(&accessToken); err != nil {
		exception := BadRequestException("error unmarshalling request json to object", err)
		context.JSON(exception.Code, exception)
		return
	}

	if exception := accessTokenWs.accessTokenService.Update(id, &accessToken); exception != nil {
		context.JSON(exception.Code, exception)
		return
	}

	context.JSON(http.StatusOK, map[string]string{})
}

func (accessTokenWs *DefaultAccessTokenWs) Delete(context *gin.Context) {

	id := context.Param("id")

	if exception := accessTokenWs.accessTokenService.Delete(id); exception != nil {
		context.JSON(exception.Code, exception)
		return
	}

	context.JSON(http.StatusOK, map[string]string{})
}

func (accessTokenWs *DefaultAccessTokenWs) FindById(context *gin.Context) {

	id := context.Param("id")

	accessToken, exception := accessTokenWs.accessTokenService.FindById(id)
	if exception != nil {
		context.JSON(exception.Code, exception)
		return
	}

	//isPublic := context.GetHeader("X-Public") == "true"
	//marshaledUser := MarshallUser(user, isPublic)

	context.JSON(http.StatusOK, accessToken)
}

func (accessTokenWs *DefaultAccessTokenWs) UpdateExpirationTime(context *gin.Context) {

}
