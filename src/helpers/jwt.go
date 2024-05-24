package helpers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type (
	ParamCreateUser struct {
		ExpiredInMinute int
		SecretKey       []byte
		UserId          int
	}
	ParamsValidateJWT struct {
		Token     string
		SecretKey string
	}

	contextKey string

	Claims struct {
		UserId int `json:"userId"`
		jwt.StandardClaims
	}
)

const UserContextKey contextKey = "user"

func CreateUserToken(p *ParamCreateUser) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserId: p.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(p.ExpiredInMinute) * time.Minute).Unix(),
		},
	})

	tokenString, err := token.SignedString(p.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func ValidateJWT(p *ParamsValidateJWT) (jwt.MapClaims, error) {
	token, err := jwt.Parse(p.Token, func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok || method != jwt.SigningMethodHS256 {
			return nil, errors.New("token tidak valid")
		}

		return []byte(p.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
