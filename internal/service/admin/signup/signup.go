package signup

import (
	"api-booking/internal/models"
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context, db *sql.DB, loginRequest SignupRequest) (SignupResponse, error) {
	response := SignupResponse{}
	// HashPassword hashes password from user input
	hashPassword, err := HashPassword(loginRequest.PassWord)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[HashPassword] error  %v", err)
		return response, err
	}
	loginRequest.PassWord = string(hashPassword)

	account := models.Account{
		Fullname:    loginRequest.FullName,
		Username:    loginRequest.UserName,
		Password:    loginRequest.PassWord,
		Email:       null.StringFrom(loginRequest.Email),
		Phonenumber: loginRequest.Phone,
		Role:        loginRequest.Role,
	}

	err = account.Insert(c, db, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[SignUp] error  %v", err)
		return response, err
	}
	response.Status = true
	response.Message = "Tạo tài khoản mới thành công"
	return response, nil
}

// HashPassword hashes password from user input
func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10) // 10 is the cost for hashing the password.
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[HashPassword] error  %v", err)
		return nil, errors.New("Đăng nhập thất bại")
	}
	return bytes, err
}
