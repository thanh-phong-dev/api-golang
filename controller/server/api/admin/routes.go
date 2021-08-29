package admin

import (
	"api-booking/controller/server/api/admin/account"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Routes(rg *gin.RouterGroup, db *sql.DB) {
	ping := rg.Group("/admin")
	accountRoutes := ping.Group("/")
	{
		account.Routes(accountRoutes, db)
	}
}
