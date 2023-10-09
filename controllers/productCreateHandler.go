package controllers

import (
  "hotelbookingapi/database"
  "hotelbookingapi/models"
  "hotelbookingapi/utils"
  "github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
  cookie := c.Cookies("jwt")
  _, err := utils.ValidateCookie(cookie)

  // check for authentication
  if err != nil {

    c.Status(fiber.StatusUnauthorized)
    return c.JSON(fiber.Map{
      "message": "Authentication faild",
      "data": nil,
      "success": false,
    })
  }
  
  var productInput models.Product 

  // parse inputs
  if err := c.BodyParser(&productInput); err != nil {
    c.Status(fiber.StatusBadRequest)
    return c.JSON(fiber.Map{
      "message": "Could not parse the data",
      "data": nil,
      "success": false,
    })
  }

  var product models.Product

  product.SetProduct(productInput)
  // validate inputs

  // create product 
  result := databse.DB.Create(&product)
  if result.Error != nil {
    c.Status(fiber.StatusUnprocessableEntity)
    return c.JSON(fiber.Map{
      "message": "Couldn't create the product",
      "data": nil,
      "success": false,
    })
  }


  return c.JSON(fiber.Map{
    "message": "Successfully created the product",
    "success": true,
    "data": product.ID,
  })
}





