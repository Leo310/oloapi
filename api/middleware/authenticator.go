package middleware

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var jwtKey = []byte(os.Getenv("PRIV_KEY"))

type claims struct {
	jwt.StandardClaims
}

// Authenticator returns a middleware which secures all the private routes
// TODO check if user with this uuid still exists. Because user could delete himself and still send private requests with his valid tokens
func Authenticator() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		claims := new(claims)

		authHeaderContent := ctx.Get("Authorization", "no_token")
		if authHeaderContent == "no_token" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(merror{ErrorCode: errTokenUnavailable})
		}
		length := len(authHeaderContent)
		// remove Bearer string at the beginning of Authorization header
		token := authHeaderContent[7:length]

		parsedToken, err := jwt.ParseWithClaims(token, claims,
			func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

		//parsedToken is nil when authorization header is empty string
		if parsedToken != nil && parsedToken.Valid {
			// expired at checked by ParseWithClaims()
			ctx.Locals("uuid", claims.Issuer)
			return ctx.Next()
		}

		ve := err.(*jwt.ValidationError)
		//filtering exact error
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			// this is not even a token, we should delete the cookies here
			return ctx.Status(fiber.StatusForbidden).JSON(merror{ErrorCode: errTokenBad})
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return ctx.Status(fiber.StatusUnauthorized).JSON(merror{ErrorCode: errTokenNotActiveExpired})
		} else {
			// cannot handle this token
			return ctx.Status(fiber.StatusForbidden).JSON(merror{ErrorCode: errTokenBad})
		}
	}
}
