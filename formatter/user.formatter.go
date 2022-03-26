package formatter

type FormatUser struct {
	Email            string `json:"email" binding:"required"`
	FirstName        string `json:"firstname"  binding:"required"`
	Lastname         string `json:"lastname" binding:"required"`
	Username         string `json:"username" binding:"required"`
	Address          string `json:"address" binding:"required"`
	Phone            uint64 `json:"phone" binding:"required"`
	Role             uint8  `json:"role"`
	Password         string `json:"password" binding:"required"`
	Confirm_Password string `json:"confirm_password" binding:"required"`
}
