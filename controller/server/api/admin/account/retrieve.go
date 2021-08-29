package account

import (
	"api-booking/internal/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAccounts(c *gin.Context, db *sql.DB) {
	accounts, err := models.Accounts().All(c, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, accounts)
}
