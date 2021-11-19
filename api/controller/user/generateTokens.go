package user

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type claims struct {
	jwt.StandardClaims
}

// generateTokens returns the access and refresh tokens
func (userenv *Userenv) generateTokens(uuid string) (string, string) {
	accessToken := generateAccessClaims(uuid, userenv.JwtKey, userenv.AccessTokenExpiryTime)
	refreshToken := generateRefreshClaims(uuid, userenv.JwtKey, userenv.RefreshTokenExpiryTime)
	return accessToken, refreshToken
}

// generateAccessClaims returns a claim and a acess_token string
func generateAccessClaims(uuid string, secret []byte, expireAt time.Duration) string {
	t := time.Now()
	claim := claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    uuid,
			ExpiresAt: t.Add(expireAt).Unix(),
			Subject:   "access_token",
			IssuedAt:  t.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}

	return tokenString
}

// generateRefreshClaims returns refresh_token
func generateRefreshClaims(uuid string, secret []byte, expireAt time.Duration) string {
	t := time.Now()
	refreshClaim := claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    uuid,
			ExpiresAt: t.Add(expireAt).Unix(),
			Subject:   "refresh_token",
			IssuedAt:  t.Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim)
	refreshTokenString, err := refreshToken.SignedString(secret)
	if err != nil {
		panic(err)
	}

	return refreshTokenString
}
