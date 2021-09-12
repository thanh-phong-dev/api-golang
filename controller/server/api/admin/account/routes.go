package account

import (
	"api-booking/middlewares"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func Routes(r *gin.RouterGroup, db *sql.DB, client *redis.Client) {
	r.POST("signup", func(c *gin.Context) { Singnup(c, db) })
	r.POST("login", func(c *gin.Context) { Login(c, db, client) })
	// router for accounts
	routerGroupAccount := r.Group("account/")
	routerGroupAccount.Use(middlewares.AuthJwtVerify())
	routerGroupAccount.Use(middlewares.CheckScopeAccess(client, middlewares.ADMIN))
	{
		routerGroupAccount.GET("views", func(c *gin.Context) { GetAccounts(c, db) })
	}
}
