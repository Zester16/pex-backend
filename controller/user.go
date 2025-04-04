package controller

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"pex.oschmid.com/helper"
	"pex.oschmid.com/model"
	"pex.oschmid.com/repository"
)

//var SecretKey = []byte("")

//var SecretKey = []byte(helper.GetJWTSecretKey())

// endpoint to create session cookie and token
func Login(c *fiber.Ctx) error {
	username := os.Getenv("username")
	password := os.Getenv("password")
	p := new(model.UserForm)
	err := c.BodyParser(p)
	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(p)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{"statusCode": 400, "error": err})
	}
	if p.Username == username && p.Password == password {
		sId := uuid.New()
		created := time.Now().UTC().Unix()
		//expires := time.Now().UTC().Unix()
		//original
		expires := time.Now().UTC().Add(time.Hour * 720)

		c.Cookie(&fiber.Cookie{
			Name:     "sid",
			Value:    sId.String(),
			MaxAge:   int(720 * time.Hour),
			Expires:  expires,
			Secure:   false,
			HTTPOnly: true,
			SameSite: "none",
		})

		sessionModel := model.SessionModel{
			Id:         sId.String(),
			Username:   username,
			Device:     string(c.Context().UserAgent()),
			Created_At: created,
			Expiry:     expires.Unix(),
		}
		err = repository.AddSession(sessionModel)

		if err != nil {
			fmt.Println("router-user-login-error", err)
			return c.JSON(&fiber.Map{"statusCode": 400, "statusMessage": "could not log you in"})
		}

		j, err := generateJWT(username, sId.String(), c)

		if len(j) == 0 {
			return err
		}

		return c.JSON(&fiber.Map{"statusCode": 200, "data": "success", "token": j})
	}

	return c.Status(401).JSON(&fiber.Map{"statusCode": 401, "error": "Forbidden"})
}

// endpoint to generate JWT token post token expiring, provided cookie is present
func RegenerateToken(c *fiber.Ctx) error {
	//t := c.Get("Authorization")
	cSid := ""
	c.Request().Header.VisitAllCookie(func(key, value []byte) {
		fmt.Println("req headerKey", string(key), "value", string(value))
		if string(key) == "sid" {
			cSid = string(value)
		}
	})
	// if len(t) == 0 {

	// 	return c.Status(403).JSON(&fiber.Map{"statusCode": 1, "statusMessage": "Un-Authorized"})
	// }

	session, err := repository.GetSession(cSid)

	if err != nil {
		fmt.Println("controller-user-regenerateToken-err", err)
		return c.Status(403).JSON(&fiber.Map{
			"statusCode": 1, "statusMessage": "Forbidden",
		})
	}

	token, err := generateJWT(session.Username, cSid, c)

	if len(token) == 0 {
		fmt.Println("controller-user-regenerateToken-err", err)
		return err
	}
	updatedSessionTime := time.Now().UTC().Unix()
	go repository.UpdateLoginTimeForSession(cSid, updatedSessionTime)
	return c.JSON(&fiber.Map{"statusCode": 0, "token": token})
}

//*********************MIDDLEWARE ********************************
//middleware to check
//func MiddlewareCheckSession() error{}

// middleware to parse jwt
func MiddlewareCheckUser(c *fiber.Ctx) error {
	t := c.Get("Authorization")
	cSid := ""
	c.Request().Header.VisitAllCookie(func(key, value []byte) {
		fmt.Println("req headerKey", string(key), "value", string(value))
		if string(key) == "sid" {
			cSid = string(value)
		}
	})
	if len(t) == 0 {
		return c.Status(403).JSON(&fiber.Map{"statusCode": 1, "statusMessage": "Forbidden"})
	}

	if len(cSid) == 0 {
		return c.Status(403).JSON(&fiber.Map{"statusCode": 2, "statusMessage": "Forbidden"})
	}

	token, err := parseToken(t, c)

	if token == nil || err != nil {
		fmt.Println(token, err)
		return err
	}
	clm := token.Claims.(jwt.MapClaims)
	//tSid = token.Raw.
	fmt.Println(clm["sid"])

	if clm["sid"] != cSid {
		fmt.Println("controller-user-MiddlewareCheck-User-err", "token cid not matching")
		return c.Status(401).JSON(&fiber.Map{"statusCode": 401, "statusMessage": "Un Authorized"})
	}

	dbResp, err := repository.GetSession(cSid)

	fmt.Println("Middleware-check-user", dbResp, err)
	if err != nil {
		return c.Status(403).JSON(&fiber.Map{"statusCode": 3, "statusMessage": "Forbidden"})
	}
	return c.Next()
}

// ******************* Auxilary or helper functions ************************************
// checks whether token is valid or not
func parseToken(t string, c *fiber.Ctx) (*jwt.Token, error) {
	SecretKey := helper.GetJWTSecretKey()
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	// Check for verification errors
	if err != nil {
		fmt.Println("controller-user-MiddlewareCheck-User-err", err)
		if strings.Contains(err.Error(), "token is malformed") {
			return nil, c.Status(403).JSON(&fiber.Map{"statusCode": 2, "statusMessage": "Token is not right"})
		}
		return nil, c.Status(401).JSON(&fiber.Map{"statusCode": 401, "statusMessage": "Token Expired"})

	}
	// Check if the token is valid
	fmt.Println(token.Valid)
	if !token.Valid {
		//fmt.Errorf("invalid token")
		return nil, c.Status(401).JSON(&fiber.Map{"statusCode": 403, "statusMessage": "Invalid token"})
	}
	return token, nil
}

// modular function to generate token
func generateJWT(username string, sId string, c *fiber.Ctx) (string, error) {
	SecretKey := helper.GetJWTSecretKey()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,                         // Subject (user identifier)
		"iss": "pex",                            // Issuer
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                // Issued at
		"sid": sId,
	})

	j, err := claims.SignedString(SecretKey)

	if err != nil {
		fmt.Println("routes-user-login-jwt-error", err)
		return "", c.Status(500).JSON(&fiber.Map{
			"statusCode":    500,
			"statusMessage": "some error happened",
		})
	}
	return j, nil

}
