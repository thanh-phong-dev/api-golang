package middlewares

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Vars struct {
	Role       string
	AccessUuid string
	UserName   string
}

// AuthJwtVerify verify token and add userID to the request context
func AuthJwtVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := extractToken(c)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("tokenString err : %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Mã xác thực không tồn tại, vui lòng đăng nhập lại"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			//Make sure that the token method conform to "SigningMethodHMAC"
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("ACCESS_SECRET")), nil
		})

		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[AuthJwtVerify] Parse token error : %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Phiên đăng nhập đã hết hạn, vui lòng đăng nhập lại"})
			c.Abort()
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Mã xác thực không hợp lệ"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || claims["Access_uuid"] == "" || claims["Role"] == "" || claims["ExpiresAt"] == "" || claims["UserID"] == "" {
			logrus.WithFields(logrus.Fields{}).Error("[AuthJwtVerify] Access_uuid,Role,ExpiresAt,UserID is null")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Mã xác thực không hợp lệ: xác thực không thành công"})
			c.Abort()
			return
		}

		accessUuid, ok := claims["Access_uuid"].(string)
		if !ok {
			logrus.WithFields(logrus.Fields{}).Error("[AuthJwtVerify] Access_uuid parse string error")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Mã xác thực không hợp lệ: xác thực không thành công"})
			c.Abort()
			return
		}

		role, ok := claims["Role"].(string)
		if !ok {
			logrus.WithFields(logrus.Fields{}).Error("[AuthJwtVerify] Role parse string error")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Mã xác thực không hợp lệ: xác thực không thành công"})
			c.Abort()
			return
		}

		userName, ok := claims["UserID"].(string)
		if !ok {
			logrus.WithFields(logrus.Fields{}).Error("[AuthJwtVerify] UserID parse string error")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Mã xác thực không hợp lệ: xác thực không thành công"})
			c.Abort()
			return
		}

		if count := getTokenRemainingValidity(claims["ExpiresAt"]); count == -1 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Mã xác thực đã hết hạn, vui lòng đăng nhập lại"})
			c.Abort()
			return
		}

		c.Request = c.Request.WithContext(context.WithValue(c, "values", Vars{Role: role, AccessUuid: accessUuid, UserName: userName})) // adding the Role to the context
		c.Next()
	}
}

func getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds())
		}
	}
	return -1
}

func extractToken(c *gin.Context) (string, error) {
	bearToken := c.GetHeader("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) != 2 {
		return "", errors.New("Mã xác thực không được cung cấp hoặc không đúng định dạng")
	}
	return strArr[1], nil
}
