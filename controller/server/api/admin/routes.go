package admin

import (
	"api-booking/controller/server/api/admin/account"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func Routes(rg *gin.RouterGroup, db *sql.DB, client *redis.Client) {
	r := rg.Group("/admin")
	accountRoutes := r.Group("/")
	{
		account.Routes(accountRoutes, db, client)
	}
}
