package login

import (
	"api-booking/internal/models"
	"api-booking/utils"
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context, db *sql.DB, client *redis.Client, loginRequest AccountRequest) (LoginReponse, error) {
	var loginReponse LoginReponse
	user, err := models.Accounts(
		models.AccountWhere.Username.EQ(loginRequest.UserName),
	).One(c, db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[GetAccountByUserName] error  %v", err)
		return loginReponse, errors.New("Đăng nhập thất bại")
	}

	if user == nil {
		return loginReponse, errors.New("Đăng nhập thất bại")
	}

	err = CheckPasswordHash(loginRequest.PassWord, user.Password)
	if err != nil {
		return loginReponse, errors.New("Đăng nhập thất bại")
	}

	token, err := utils.EncodeAuthToken(loginRequest.UserName, user.Role)
	if err != nil {
		return loginReponse, err
	}

	saveErr := utils.CreateAuth(token, client)
	if saveErr != nil {
		return loginReponse, saveErr
	}

	loginReponse.Success = true
	loginReponse.Token = token.AccessToken
	loginReponse.InfoUser = User{
		FullName: user.Fullname,
		UserName: user.Username,
		Email:    user.Email.String,
		Phone:    user.Phonenumber,
		Role:     user.Role,
	}

	return loginReponse, nil
}

// CheckPasswordHash checks password hash and password from user input if they match
func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CheckPasswordHash] error  %v", err)
		return errors.New("Đăng nhập thất bại")
	}
	return nil
}
