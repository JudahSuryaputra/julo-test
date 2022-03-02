package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func EncodeAuthToken(id uuid.UUID) (string, error) {
	claims := jwt.MapClaims{}
	claims["ID"] = id
	claims["CreatedAt"] = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	return token.SignedString([]byte(viper.GetString("SECRET")))
}
