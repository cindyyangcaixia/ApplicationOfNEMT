package utils

import (
	"strings"
	"time"

	"github.com/cindyyangcaixia/ApplicationOfNEMT/pkg/setting"
	"github.com/golang-jwt/jwt/v5"
)

type Schema string

const (
	Client Schema = "client"
	Admin  Schema = "admin"
)

type CustomClaims struct {
	UserID int `json:"userId"`
	jwt.RegisteredClaims
}

// func authValues() map[string]bool {
// 	return map[string]bool{
// 		string(Client): true,
// 		string(Admin):  true,
// 	}
// }

// func isAuthValue(schema Schema) bool {
// 	return authValues()[string(schema)]
// }

func GenerateToken(UserID int, schema string) (string, error) {

	claims := CustomClaims{
		UserID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "ApplicationOfNEMT",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(getJwtSecret(schema)))

	return tokenString, err
}

func GetToken(authorization string) string {
	if !strings.Contains(authorization, "Bearer") {
		return ""
	}

	token := strings.Split(authorization, "Bearer ")
	return token[1]
}

func getJwtSecret(schema string) string {
	switch schema {
	case string(Client):
		return setting.AppSetting.JwtSecret
	case string(Admin):
		return setting.AppSetting.JwtSecret
	default:
		return setting.AppSetting.JwtSecret
	}
}

func VerifyToken(token string, schema string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(getJwtSecret(schema)), nil
	})

	if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, err
}
