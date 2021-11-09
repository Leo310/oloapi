package user

import (
	db "oloapi/api/database"
	"oloapi/api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// generateTokens returns the access and refresh tokens
func generateTokens(uuid string) (string, string) {
	claim, accessToken := generateAccessClaims(uuid)
	refreshToken := generateRefreshClaims(claim)

	return accessToken, refreshToken
}

// generateAccessClaims returns a claim and a acess_token string
func generateAccessClaims(uuid string) (*models.Claims, string) {
	t := time.Now()
	claim := &models.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    uuid,
			ExpiresAt: t.Add(15 * time.Minute).Unix(),
			Subject:   "access_token",
			IssuedAt:  t.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return claim, tokenString
}

// generateRefreshClaims returns refresh_token
func generateRefreshClaims(cl *models.Claims) string {
	result := db.DB.Where(&models.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer: cl.Issuer,
		},
	}).Find(&models.Claims{})

	// checking the number of refresh tokens stored.
	// If the number is higher than 3, remove all the refresh tokens and leave only new one.
	if result.RowsAffected > 3 {
		db.DB.Where(&models.Claims{
			StandardClaims: jwt.StandardClaims{Issuer: cl.Issuer},
		}).Delete(&models.Claims{})
	}

	t := time.Now()
	refreshClaim := &models.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    cl.Issuer,
			ExpiresAt: t.Add(30 * 24 * time.Hour).Unix(),
			Subject:   "refresh_token",
			IssuedAt:  t.Unix(),
		},
	}

	// create a claim on DB
	db.DB.Create(&refreshClaim)

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim)
	refreshTokenString, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return refreshTokenString
}

// getAuthCookies sends two cookies of type access_token and refresh_token
// TODO turn secure back on when https. postman couldnt recognize these cookies. was such a pain in the ass
func getAuthCookies(accessToken, refreshToken string) (*fiber.Cookie, *fiber.Cookie) {
	accessCookie := &fiber.Cookie{
		Name:    "access_token",
		Value:   accessToken,
		Expires: time.Now().Add(24 * time.Hour),
		// HTTPOnly: true,
		// Secure:   true,
	}

	refreshCookie := &fiber.Cookie{
		Name:    "refresh_token",
		Value:   refreshToken,
		Expires: time.Now().Add(10 * 24 * time.Hour),
		// HTTPOnly: true,
		// Secure:   true,
	}

	return accessCookie, refreshCookie
}
