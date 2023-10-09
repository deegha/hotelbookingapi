package controllers
import (
  "hotelbookingapi/database"
  "hotelbookingapi/models"
  "hotelbookingapi/utils"
  "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UpdateProduct(c *fiber.Ctx) error {
  cookie := c.Cookies("jwt")
  _, err := utils.ValidateCookie(cookie)

  // check for authentication
  if err != nil {
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

  productId := c.Params("id")

  var existingProduct models.Product
  if err := databse.DB.Preload("Media").First(&existingProduct, productId).Error; err != nil {
    return err
  }

  existingProduct.SetProduct(productInput)
  // validate inputs

  // update product 
  result := databse.DB.Session(&gorm.Session{FullSaveAssociations: false}).Save(&existingProduct)
  if result.Error != nil {
    c.Status(fiber.StatusUnprocessableEntity)
    return c.JSON(fiber.Map{
      "message": "Couldn't update the product",
      "data": nil,
      "success": false,
    })
  }

  return c.JSON(fiber.Map{
    "message": "Successfully updated the product",
    "success": true,
    "data": nil,
  })
}
