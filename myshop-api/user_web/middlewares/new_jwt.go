package middlewares

import (
	"myshop-api/user_web/global"
	// "myshop-api/user_web/models"

	// "github.com/golang-jwt/jwt/v5"
)

type new_jwjt struct {
	SignKey []byte
}

func New() *new_jwjt{
	return &new_jwjt{
		[]byte(global.UserWebInfo.JwtInfo.Key),
	}
}

// func (j *JWT)CreateJwtToken(claims models.CustomClaims) (string, error) {
// 	// Create a new token object, specifying signing method and the claims
// 	// you would like it to contain.
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	// Sign and get the complete encoded token as a string using the secret
// 	return token.SignedString(j.SigningKey)
// }