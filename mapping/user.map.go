package mapping

import "triadmoko-be-golang/entity"

type FormatUser struct {
	Email            string `json:"email" binding:"required, email"`
	FirstName        string `json:"firstname"  binding:"required"`
	Lastname         string `json:"lastname" binding:"required"`
	Username         string `json:"username" binding:"required"`
	Address          string `json:"address" binding:"required"`
	Phone            uint64 `json:"phone" binding:"required"`
	Role             uint8  `json:"role"`
	Password         string `json:"password" binding:"required"`
	Confirm_Password string `json:"confirm_password" binding:"required"`
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
