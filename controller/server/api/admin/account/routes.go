package account

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, db *sql.DB) {
	r.GET("accounts", func(c *gin.Context) {
		GetAccounts(c, db)
	})
}
