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
	signupResponse := SignupResponse{}

	isUserNameUnique, err := checkUserName(c, db, loginRequest.UserName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CheckUserName] error  %v", err)
		return signupResponse, errors.New("Lỗi hệ thống")
	}

	if isUserNameUnique {
		return signupResponse, errors.New("Tên đăng nhập đã tồn tại")
	}

	if loginRequest.Email != "" {
		isEmailUnique, err := checkEmail(c, db, loginRequest.Email)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[CheckUserName] error  %v", err)
			return signupResponse, errors.New("Lỗi hệ thống")
		}

		if isEmailUnique {
			return signupResponse, errors.New("Email đã tồn tại")
		}

	}

	// HashPassword hashes password from user input
	hashPassword, err := HashPassword(loginRequest.PassWord)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[HashPassword] error  %v", err)
		return signupResponse, err
	}
	loginRequest.PassWord = string(hashPassword)

	account := models.Account{
		Fullname:    loginRequest.FullName,
		Username:    loginRequest.UserName,
		Password:    loginRequest.PassWord,
		Email:       null.StringFrom(loginRequest.Email),
		Phonenumber: loginRequest.Phone,
		Role:        loginRequest.Role,
		Sex:         loginRequest.Sex,
		DateOfBirth: null.StringFrom(loginRequest.DateOfBirth),
		Address:     null.StringFrom(loginRequest.Address),
	}

	err = account.Insert(c, db, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[SignUp] error  %v", err)
		return signupResponse, errors.New("Lỗi hệ thống")
	}
	signupResponse.Status = true
	signupResponse.Message = "Tạo tài khoản mới thành công"
	return signupResponse, nil
}

func checkUserName(c *gin.Context, db *sql.DB, userName string) (bool, error) {
	account, err := models.Accounts(
		models.AccountWhere.Username.EQ(userName),
	).Exists(c, db)
	if err != nil {
		return false, err
	}
	return account, nil
}

func checkEmail(c *gin.Context, db *sql.DB, email string) (bool, error) {
	account, err := models.Accounts(
		models.AccountWhere.Email.EQ(null.StringFrom(email)),
	).Exists(c, db)
	if err != nil {
		return false, err
	}
	return account, nil
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
