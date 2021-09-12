package signup

type SignupResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type SignupRequest struct {
	Email       string `json:"email"`
	UserName    string `json:"username"`
	PassWord    string `json:"password"`
	Role        string `json:"role"`
	Sex         string `json:"sex"`
	DateOfBirth string `json:"dateofbirth"`
	Phone       string `json:"phone"`
	FullName    string `json:"fullname"`
	Address     string `json:"address"`
}
