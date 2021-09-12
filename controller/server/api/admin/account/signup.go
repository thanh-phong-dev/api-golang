package account

import (
	"api-booking/internal/service/admin/signup"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Singnup(c *gin.Context, db *sql.DB) {
	var signupRequest signup.SignupRequest
	if err := c.ShouldBindJSON(&signupRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	login, err := signup.SignUp(c, db, signupRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, login)
}
