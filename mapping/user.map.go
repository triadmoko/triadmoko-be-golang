package mapping

import "triadmoko-be-golang/entity"

type FormatUser struct {
	Email            string `json:"email"`
	FirstName        string `json:"firstname" `
	Lastname         string `json:"lastname"`
	Username         string `json:"username"`
	Address          string `json:"address"`
	Phone            uint64 `json:"phone"`
	Role             uint8  `json:"role"`
	Password         string `json:"password"`
	Confirm_Password string `json:"confirm_password"`
}
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func UserFormatter(user entity.User) FormatUser {
	format := FormatUser{
		Email:     user.Email,
		FirstName: user.Firstname,
		Lastname:  user.Lastname,
		Address:   user.Address,
		Phone:     user.Phone,
		Password:  user.Password,
		Role:      user.Role,
	}
	return format
}
