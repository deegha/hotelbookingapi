package controllers

import (
	"fmt"
	"hotelbookingapi/database"
	"hotelbookingapi/models"
	"hotelbookingapi/utils"
	"github.com/gofiber/fiber/v2"
)

func BookingUpdateHandler(c *fiber.Ctx) error {
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

  var bookingInput models.Booking 

  // parse inputs
  if err := c.BodyParser(&bookingInput); err != nil {
		fmt.Print(err)
    c.Status(fiber.StatusBadRequest)
    return c.JSON(fiber.Map{
      "message": "Could not parse the data",
      "data": nil,
      "success": false,
    })
  }

  bookingId := c.Params("id")
	var existingBooking models.Booking
  if err := databse.DB.First(&existingBooking, bookingId).Error; err != nil {
    return err
  }

  existingBooking.SetBooking(bookingInput)

	result := databse.DB.Save(&existingBooking)
	if result.Error != nil {
		fmt.Print(err)
    c.Status(fiber.StatusUnprocessableEntity)
    return c.JSON(fiber.Map{
      "message": "Couldn't create the booking",
      "data": nil,
      "success": false,
    })
  }

  bookingReturn := BookingReturn {
		BookingID: int(existingBooking.ID),
	}

  return c.JSON(fiber.Map{
    "message": "Successfully updated the booking",
    "success": true,
    "data": bookingReturn,
  })
}


