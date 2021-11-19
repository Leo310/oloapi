package user

import (
	"oloapi/api/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Authenticator returns a middleware which secures all the private routes
func (userenv *Userenv) Authenticator() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		claims := new(claims)

		authHeaderContent := ctx.Get("Authorization", "no_token")
		if authHeaderContent == "no_token" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(uerror{ErrorCode: errTokenUnavailable})
		}
		length := len(authHeaderContent)
		// remove Bearer string at the beginning of Authorization header
		token := authHeaderContent[7:length]

		parsedToken, err := jwt.ParseWithClaims(token, claims,
			func(token *jwt.Token) (interface{}, error) {
				return userenv.JwtKey, nil
			})

		//parsedToken is nil when authorization header is empty string
		if parsedToken != nil && parsedToken.Valid {
			// check if user still exsits in db
			if count := userenv.DB.Where("uuid = ?", claims.Issuer).First(&models.User{}).RowsAffected; count == 0 {
				return ctx.Status(fiber.StatusForbidden).JSON(uerror{ErrorCode: errTokenOfNonexistingUser})
			}
			// expired at checked by ParseWithClaims()
			ctx.Locals("uuid", claims.Issuer)
			return ctx.Next()
		}

		ve := err.(*jwt.ValidationError)
		//filtering exact error
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			// this is not even a token, we should delete the cookies here
			return ctx.Status(fiber.StatusForbidden).JSON(uerror{ErrorCode: errTokenBad})
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return ctx.Status(fiber.StatusUnauthorized).JSON(uerror{ErrorCode: errTokenNotActiveExpired})
		} else {
			// cannot handle this token
			return ctx.Status(fiber.StatusForbidden).JSON(uerror{ErrorCode: errTokenBad})
		}
	}
}
