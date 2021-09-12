package login

type AccountRequest struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type LoginReponse struct {
	Success  bool   `json:"success"`
	Token    string `json:"token"`
	InfoUser User   `json:"infoUser"`
}

type User struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	FullName string `json:"fullname"`
}
