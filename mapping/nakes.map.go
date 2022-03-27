package mapping

type InputNakes struct {
	NoNakes  int    `json:"no_nakes" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Addres   string `json:"address" binding:"required"`
	FaskesID int    `json:"faskes_id" binding:"required"`
}
