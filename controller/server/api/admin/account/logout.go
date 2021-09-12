package account

import (
	"api-booking/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// LogoutAccount controller for logout
func LogoutAccount(c *gin.Context, client *redis.Client) {
	claims, err := middlewares.ClaimsToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Mã xác thực không tồn tại, vui lòng đăng nhập lại"})
		return
	}
	deleted, delErr := DeleteAuth(claims.AccessUuid, client)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Phiên đăng nhập đã hết hạn, vui lòng đăng nhập lại"})
		return
	}
	// send Result response
	c.JSON(http.StatusOK, gin.H{"Message": "Đã đăng xuất"})
}

func DeleteAuth(givenUuid string, client *redis.Client) (int64, error) {
	deleted, err := client.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
