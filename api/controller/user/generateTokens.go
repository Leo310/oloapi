package user

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	accessTokenTime  time.Duration = 15 * time.Minute
	refreshTokenTime time.Duration = 30 * 24 * time.Hour
)

type claims struct {
	jwt.StandardClaims
}

// generateTokens returns the access and refresh tokens
func generateTokens(uuid string) (string, string) {
	accessToken := generateAccessClaims(uuid)
	refreshToken := generateRefreshClaims(uuid)

	return accessToken, refreshToken
}

// generateAccessClaims returns a claim and a acess_token string
func generateAccessClaims(uuid string) string {
	t := time.Now()
	claim := claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    uuid,
			ExpiresAt: t.Add(accessTokenTime).Unix(),
			Subject:   "access_token",
			IssuedAt:  t.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return tokenString
}

// generateRefreshClaims returns refresh_token
func generateRefreshClaims(uuid string) string {
	t := time.Now()
	refreshClaim := claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    uuid,
			ExpiresAt: t.Add(refreshTokenTime).Unix(),
			Subject:   "refresh_token",
			IssuedAt:  t.Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim)
	refreshTokenString, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return refreshTokenString
}
