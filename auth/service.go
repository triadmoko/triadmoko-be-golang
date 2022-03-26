package auth

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(ID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
type jwtService struct {
}

var SECRET_KEY_USER = []byte(os.Getenv("JWT_SECRET"))

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(ID int) (string, error) {
	payload := jwt.MapClaims{}
	payload["user_id"] = ID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, err := token.SignedString(SECRET_KEY_USER)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}
func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, oke := token.Method.(*jwt.SigningMethodHMAC)
		if !oke {
			return nil, errors.New("Invalid Token")
		}
		return []byte(SECRET_KEY_USER), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil

}
