package formatter

import (
	"triadmoko-be-golang/entity"
)

type FormatUser struct {
	Email            string `json:"email" binding:"required,email"`
	FirstName        string `json:"firstname"  binding:"required"`
	Lastname         string `json:"lastname" binding:"required"`
	Username         string `json:"username" binding:"required"`
	Address          string `json:"address" binding:"required"`
	Phone            uint64 `json:"phone" binding:"required"`
	Role             uint8  `json:"role"`
	Password         string `json:"password" binding:"required"`
	Confirm_Password string `json:"confirm_password" binding:"required"`
}

type ResponseRegisterUser struct {
	Email     string `json:"email"`
	FirstName string `json:"firstname" `
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Address   string `json:"address"`
	Phone     uint64 `json:"phone"`
}
type UserFormatter struct {
	Token string `json:"token"`
}

func UserJsonFormatter(token string) UserFormatter {
	formatter := UserFormatter{
		Token: token,
	}
	return formatter
}
func FormateResponseRegisterUser(user entity.User) ResponseRegisterUser {
	response := ResponseRegisterUser{
		Email:     user.Email,
		FirstName: user.Firstname,
		Lastname:  user.Lastname,
		Username:  user.Username,
		Address:   user.Address,
		Phone:     user.Phone,
	}
	return response
}
