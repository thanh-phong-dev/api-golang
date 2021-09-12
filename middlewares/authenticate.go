package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

const (
	ADMIN = "ADMIN"
)

func CheckScopeAccess(client *redis.Client, scopes ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := ClaimsToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Mã xác thực không tồn tại, vui lòng đăng nhập lại"})
			c.Abort()
			return
		}
		uuid, err := FetchAuth(claims.AccessUuid, client)
		if err != nil || len(uuid) == 0 {
			logrus.WithFields(logrus.Fields{}).Error("FetchAuth error", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Phiên đăng nhập đã hết hạn, vui lòng đăng nhập lại"})
			c.Abort()
			return
		}
		logrus.WithFields(logrus.Fields{}).Infof("User %s Role %s loging ", claims.UserName, claims.Role)
		for _, scope := range scopes {
			if claims.Role != scope {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Bạn không có quyền truy cập"})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

func FetchAuth(givenUuid string, client *redis.Client) (string, error) {
	userid, err := client.Get(givenUuid).Result()
	if err != nil {
		return "", err
	}
	return userid, nil
}
