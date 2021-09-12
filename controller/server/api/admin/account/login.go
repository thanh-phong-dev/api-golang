package account

import (
	"api-booking/internal/service/admin/login"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func Login(c *gin.Context, db *sql.DB, client *redis.Client) {
	var accountRequest login.AccountRequest
	if err := c.ShouldBindJSON(&accountRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate(accountRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	login, err := login.Login(c, db, client, accountRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, login)
}

func validate(accountRequest login.AccountRequest) error {
	if accountRequest.UserName == "" {
		return errors.New("Bạn chưa nhập tên đăng nhập")
	}

	if len(accountRequest.UserName) > 100 {
		return errors.New("Tên đăng nhập quá dài")
	}

	if accountRequest.PassWord == "" {
		return errors.New("Bạn chưa nhập mật khẩu")
	}

	if len(accountRequest.PassWord) > 100 {
		return errors.New("Mật khẩu quá dài")
	}

	return nil
}
