package middleware

import (
	"log"
	"oloapi/api/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var jwtKey = []byte(os.Getenv("PRIV_KEY"))

// Authenticator returns a middleware which secures all the private routes
// TODO check if user with this uuid still exists. Because user could delete himself and still send private requests with his valid tokens
func Authenticator() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		accessToken := c.Cookies("access_token", "no_token")
		claims := new(models.Claims)

		if accessToken == "no_token" {
			log.Println(`Couldn't find "access_token" cookie. Checking Authorization header.`)
			authHeaderContent := c.Get("Authorization", "no_token")
			length := len(authHeaderContent)
			accessToken = authHeaderContent[7:length]
		}

		if accessToken == "no_token" {
			return c.Status(fiber.StatusUnauthorized).JSON(merror{ErrorCode: errTokenUnavailable})
		}

		token, err := jwt.ParseWithClaims(accessToken, claims,
			func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

		if token != nil && token.Valid {
			if claims.ExpiresAt < time.Now().Unix() {
				return c.Status(fiber.StatusUnauthorized).JSON(merror{ErrorCode: errTokenExpired})
			}
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// this is not even a token, we should delete the cookies here
				c.ClearCookie("access_token", "refresh_token")
				return c.SendStatus(fiber.StatusForbidden)
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return c.Status(fiber.StatusUnauthorized).JSON(merror{ErrorCode: errTokenNotActiveExpired})
			} else {
				// cannot handle this token
				c.ClearCookie("access_token", "refresh_token")
				return c.SendStatus(fiber.StatusForbidden)
			}
		}

		c.Locals("uuid", claims.Issuer)
		return c.Next()
	}
}
