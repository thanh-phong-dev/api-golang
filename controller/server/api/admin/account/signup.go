package account

import (
	"api-booking/internal/response"
	"api-booking/internal/service/admin/signup"
	"api-booking/utils"
	"database/sql"
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
)

const (
	ADMIN string = "ADMIN"
	USER  string = "USER"
)

func Singnup(c *gin.Context, db *sql.DB) {
	var signupRequest signup.SignupRequest

	if err := c.ShouldBindJSON(&signupRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateSignUpRequest(signupRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	login, err := signup.SignUp(c, db, signupRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, login)
}

func validateSignUpRequest(signupRequest signup.SignupRequest) error {
	if signupRequest.UserName == "" {
		return &response.RequestError{Message: "Bạn chưa nhập tên đăng nhập"}
	}
	if len(signupRequest.UserName) > 100 {
		return &response.RequestError{Message: "Tên đăng nhập không được dài quá 100 ký tự"}
	}

	if signupRequest.PassWord == "" {
		return &response.RequestError{Message: "Bạn chưa nhập mật khẩu"}
	}
	if len(signupRequest.PassWord) > 100 {
		return &response.RequestError{Message: "Mật khẩu không được dài quá 100 ký tự"}
	}
	if err := utils.VerifyPassword(signupRequest.PassWord); err != nil {
		return err
	}

	if signupRequest.Email != "" {
		if err := checkmail.ValidateFormat(signupRequest.Email); err != nil {
			return &response.RequestError{Message: "Email không đúng định dạng"}
		}
	}

	if signupRequest.Role == "" {
		return &response.RequestError{Message: "Bạn chưa nhập quyền"}
	}
	if signupRequest.Role != USER {
		return &response.RequestError{Message: "Phân quền không tồn tại"}
	}

	if signupRequest.Sex == "" {
		return &response.RequestError{Message: "Bạn chưa nhập giới tính"}
	}
	if len(signupRequest.Sex) > 3 {
		return &response.RequestError{Message: "Giới tính không đúng định dạng"}
	}

	if signupRequest.DateOfBirth == "" {
		return &response.RequestError{Message: "Bạn chưa nhập ngày tháng năm sinh"}
	} else {
		if len(signupRequest.DateOfBirth) > 10 || !utils.CheckDate(signupRequest.DateOfBirth) {
			return &response.RequestError{Message: "Ngày tháng năm sinh không đúng định dạng"}
		}
	}

	if signupRequest.FullName == "" {
		return &response.RequestError{Message: "Bạn chưa nhập họ và tên"}
	}
	if len(signupRequest.FullName) > 250 {
		return &response.RequestError{Message: "Họ và tên không được dài quá 250 ký tự"}
	}

	if signupRequest.Address == "" {
		return &response.RequestError{Message: "Bạn chưa nhập địa chỉ"}
	}
	if len(signupRequest.Address) > 500 {
		return &response.RequestError{Message: "Địa chỉ không được dài quá 500 ký tự"}
	}

	if signupRequest.Phone == "" {
		return &response.RequestError{Message: "Bạn chưa nhập số điện thoại"}
	} else {
		if len(signupRequest.Phone) > 15 || !utils.CheckNumber(signupRequest.Phone) {
			return &response.RequestError{Message: "Số điện thoại không đúng định dạng"}
		}
	}

	return nil
}
