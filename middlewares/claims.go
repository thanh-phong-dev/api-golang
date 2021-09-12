package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func ClaimsToken(c *gin.Context) (Vars, error) {
	resp, ok := c.Request.Context().Value("values").(Vars)
	if !ok {
		return resp, errors.New("Mã xác thực không hợp lệ: xác thực không thành công")
	}
	
	if resp.AccessUuid == "" || resp.UserName == "" || resp.Role == "" {
		return resp, errors.New("Mã xác thực không hợp lệ: xác thực không thành công")
	}

	return resp, nil
}
