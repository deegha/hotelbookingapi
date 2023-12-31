package controllers

import (
  "fmt"
  "strconv"
  "time"
  "hotelbookingapi/models" 
  "hotelbookingapi/database"
  "hotelbookingapi/utils"
  "github.com/gofiber/fiber/v2"
  "github.com/golang-jwt/jwt"
  "golang.org/x/crypto/bcrypt"
)

/*
Register function
*/
func Register(c *fiber.Ctx) error {
  var data map[string]string

  if err := c.BodyParser(&data); err != nil {
    return err
  }

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Type:     data["type"],
		Password: password,
	}

	result := databse.DB.Create(&user)

	if result.Error != nil {
		fmt.Println("Login called", result.Error)
		return c.JSON(fiber.Map{
			"message": "Could not create user, check logs",
			"success": false,
		})
	}

  return c.JSON(fiber.Map{
		"message": "Successfully created user",
		"data": c.JSON(user),
		"success": true,
	})
}

/*
Login function
*/
func Login(c *fiber.Ctx) error {
	fmt.Println("Login called")
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		fmt.Println(err)

		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Couldnt process the request body",
			"success": false,
		})
	}

	var user models.User

	databse.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
			"success": false,
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect email",
			"success": false,
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(utils.SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "couldnt sign user in",
			"success": false ,
		})
	}

	cookie := fiber.Cookie{

		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Sucessfully logged in",
		"success": true,
	})
}

/*
Logout function
*/
func Logout(c *fiber.Ctx) error {

	cookie := fiber.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Successfully logged out",
		"success": true,
	})
}
