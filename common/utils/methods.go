package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

func GenerateVerificationCode() string {
	str, verificationCode := "0123456789", ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < VerificationCodeLength; i++ {
		verificationCode += fmt.Sprintf("%c", str[rand.Intn(10)])
	}
	return verificationCode
}

func GenerateUUID() string {
	return uuid.NewV4().String()
}

func GenerateJwtToken(secreKey string, iat, seconds int64, userIdentity string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userIdentity"] = userIdentity
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secreKey))
}
