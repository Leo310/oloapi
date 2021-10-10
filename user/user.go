package user

import (
	"log"
	"math/rand"
	db "oloapi/database"
	"oloapi/models"
	"oloapi/util"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var jwtKey = []byte(os.Getenv("PRIV_KEY"))

// CreateUser route registers a User into the database
func CreateUser(c *fiber.Ctx) error {
	u := new(models.User)

	if err := c.BodyParser(u); err != nil {
		return c.JSON(fiber.Map{
			"error": true,
			"input": "Please review your input",
		})
	}

	// validate if the email, username and password are in correct format
	errors := util.ValidateRegister(u)
	if errors.Err {
		return c.JSON(errors)
	}

	if count := db.DB.Where(&models.User{Email: u.Email}).First(new(models.User)).RowsAffected; count > 0 {
		errors.Err, errors.Email = true, "Email is already registered"
	}
	if errors.Err {
		return c.JSON(errors)
	}

	// Hashing the password with a random salt
	password := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(
		password,
		rand.Intn(bcrypt.MaxCost-bcrypt.MinCost)+bcrypt.MinCost,
	)

	if err != nil {
		panic(err)
	}
	u.Password = string(hashedPassword)
	// u.ProfileImage = "https://avatars.dicebear.com/api/micah/" + u.Email + ".svg"

	if err := db.DB.Create(&u).Error; err != nil {
		return c.JSON(fiber.Map{
			"error":   true,
			"general": "Something went wrong, please try again later. 😕",
		})
	}

	// setting up the authorization cookies
	accessToken, refreshToken := util.GenerateTokens(u.UUID.String())
	accessCookie, refreshCookie := util.GetAuthCookies(accessToken, refreshToken)
	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func LoginUser(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}

	input := new(LoginInput)

	if err := c.BodyParser(input); err != nil {
		return c.JSON(fiber.Map{"error": true, "input": "Please review your input"})
	}

	// check if a user exists
	u := new(models.User)
	if res := db.DB.Where(
		&models.User{Email: input.Identity}).First(&u); res.RowsAffected <= 0 {
		return c.JSON(fiber.Map{"error": true, "general": "Invalid Credentials."})
	}

	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password)); err != nil {
		return c.JSON(fiber.Map{"error": true, "general": "Invalid Credentials."})
	}

	// setting up the authorization cookies
	accessToken, refreshToken := util.GenerateTokens(u.UUID.String())
	accessCookie, refreshCookie := util.GetAuthCookies(accessToken, refreshToken)
	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func GetUserData(c *fiber.Ctx) error {
	id := c.Locals("id")
	if c.Params("id") != "" {
		id = c.Params("id")
	}

	u := new(models.User)
	if res := db.DB.Where("uuid = ?", id).First(&u); res.RowsAffected <= 0 {
		return c.JSON(fiber.Map{"error": true, "general": "Cannot find the User"})
	}

	db.DB.Model(&u).Omit("Follows.*").Association("Followers").Count()

	return c.JSON(fiber.Map{
		"uuid":            u.UUID,
		"email":           u.Email,
		"name":            u.Name,
		"profile_picture": u.ProfileImage,
		"is_official":     u.IsOfficial,
		"followers":       db.DB.Model(&u).Omit("Follows.*").Association("Followers").Count(),
		"follows":         db.DB.Model(&u).Omit("Followers.*").Association("Follows").Count(),
	})

}

func FollowUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == c.Locals("id") {
		return c.JSON(fiber.Map{"error": true, "general": "Cannot follow your self"})
	}

	follower := new(models.User)
	if res := db.DB.Where("uuid = ?", c.Locals("id")).First(&follower); res.RowsAffected <= 0 {
		return c.JSON(fiber.Map{"error": true, "general": "Cannot find the Follower"})
	}

	u := new(models.User)
	if res := db.DB.Where("uuid = ?", id).First(&u); res.RowsAffected <= 0 {
		return c.JSON(fiber.Map{"error": true, "general": "Cannot find the User"})
	}

	err := db.DB.Model(&u).Omit("Followers.*").Association("Followers").Append(follower)
	if err != nil {
		return c.JSON(fiber.Map{
			"error":   true,
			"general": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"general": follower.Name + " follows " + u.Name + ".",
	})
}

func GetAccessToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token", "no_token")

	if refreshToken == "no_token" {
		log.Println(`Couldn't find "refresh_token" cookie. Checking Authorization header.`)
		authHeaderContent := c.Get("Authorization", "no_token")
		length := len(authHeaderContent)
		refreshToken = authHeaderContent[7:length]
	}

	refreshClaims := new(models.Claims)
	token, _ := jwt.ParseWithClaims(refreshToken, refreshClaims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if res := db.DB.Where(
		"expires_at = ? AND issued_at = ? AND issuer = ?",
		refreshClaims.ExpiresAt, refreshClaims.IssuedAt, refreshClaims.Issuer,
	).First(&models.Claims{}); res.RowsAffected <= 0 {
		// no such refresh token exist in the database
		c.ClearCookie("access_token", "refresh_token")
		return c.SendStatus(fiber.StatusForbidden)
	}

	if token.Valid {
		if refreshClaims.ExpiresAt < time.Now().Unix() {
			// refresh token is expired
			c.ClearCookie("access_token", "refresh_token")
			return c.SendStatus(fiber.StatusForbidden)
		}
	} else {
		// malformed refresh token
		c.ClearCookie("access_token", "refresh_token")
		return c.SendStatus(fiber.StatusForbidden)
	}

	_, accessToken := util.GenerateAccessClaims(refreshClaims.Issuer)

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	})

	return c.JSON(fiber.Map{"access_token": accessToken})
}
