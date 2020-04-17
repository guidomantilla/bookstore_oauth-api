package app

import (
	"strings"

	"github.com/gin-gonic/gin"
	. "github.com/guidomantilla/bookstore_common-lib/common/config"
	. "github.com/guidomantilla/bookstore_common-lib/common/nosql/impl"
	. "github.com/guidomantilla/bookstore_oauth-api/core/repository/impl"
	. "github.com/guidomantilla/bookstore_oauth-api/core/service/impl"
	. "github.com/guidomantilla/bookstore_oauth-api/core/ws"
	. "github.com/guidomantilla/bookstore_oauth-api/core/ws/impl"
)

const (
	BOOKSTORE_OAUTH_DATASOURCE_URL      = "BOOKSTORE_OAUTH_DATASOURCE_URL"
	BOOKSTORE_OAUTH_DATASOURCE_USERNAME = "BOOKSTORE_OAUTH_DATASOURCE_USERNAME"
	BOOKSTORE_OAUTH_DATASOURCE_PASSWORD = "BOOKSTORE_OAUTH_DATASOURCE_PASSWORD"
	BOOKSTORE_OAUTH_DATASOURCE_KEYSPACE = "BOOKSTORE_OAUTH_DATASOURCE_KEYSPACE"
	BOOKSTORE_OAUTH_ENVIRONMENT         = "BOOKSTORE_OAUTH_ENVIRONMENT"
)

func Init() {

	Config()

	accessTokenWs := Wire()

	ZapLogger.Info("starting the app")

	router := gin.Default()

	router.POST("/oauth/access_token", accessTokenWs.Create)
	router.GET("/oauth/access_token", accessTokenWs.Find)
	router.PUT("/oauth/access_token/:id", accessTokenWs.Update)
	router.PATCH("/oauth/access_token/:id", accessTokenWs.UpdateExpirationTime)
	router.DELETE("/oauth/access_token/:id", accessTokenWs.Delete)
	router.GET("/oauth/access_token/:id", accessTokenWs.FindById)

	router.Run(":8081")
}

func Config() {

	ConfigProperties()

	Properties[BOOKSTORE_OAUTH_DATASOURCE_URL] = "127.0.0.1"
	Properties[BOOKSTORE_OAUTH_DATASOURCE_USERNAME] = ""
	Properties[BOOKSTORE_OAUTH_DATASOURCE_PASSWORD] = ""
	Properties[BOOKSTORE_OAUTH_DATASOURCE_KEYSPACE] = "oauth"
	Properties[BOOKSTORE_OAUTH_ENVIRONMENT] = "dev"

	ConfigZapLogger(Properties[BOOKSTORE_OAUTH_ENVIRONMENT])
}

func Wire() AccessTokenWs {

	urls := strings.Split(Properties[BOOKSTORE_OAUTH_DATASOURCE_URL], ";")
	username := Properties[BOOKSTORE_OAUTH_DATASOURCE_USERNAME]
	password := Properties[BOOKSTORE_OAUTH_DATASOURCE_PASSWORD]
	keyspace := Properties[BOOKSTORE_OAUTH_DATASOURCE_KEYSPACE]

	cassandraDataSource := NewDefaultCassandraDataSource(username, password, keyspace, urls...)
	accessTokenRepository := NewDefaultAccessTokenRepository(cassandraDataSource)
	accessTokenService := NewDefaultAccessTokenService(accessTokenRepository)
	accessTokenWs := NewDefaultAccessTokenWs(accessTokenService)

	return accessTokenWs
}
