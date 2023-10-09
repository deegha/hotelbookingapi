package controllers

import (
	"fmt"
	"hotelbookingapi/database"
	"hotelbookingapi/models"
	"hotelbookingapi/utils"
	"github.com/gofiber/fiber/v2"
)

type BookingReturn struct {
	BookingID int `json:"bookingId"`
}

func BookingCreateHandler(c *fiber.Ctx) error {
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

	var booking models.Booking
	//check if the booking dates are avilable for the product 
	 checkBookingResult := databse.DB.Raw(
		"SELECT product_id FROM bookings WHERE ? < check_out_time AND ? > check_in_time AND product_id = ?", 
		bookingInput.CheckOutTime, 
	  bookingInput.CheckInTime,
		bookingInput.ProductID).Scan(&booking)

  if checkBookingResult.Error !=  nil {
		return c.JSON(fiber.Map{
      "message": "Sorry there is a booking at that time in the product",
      "data": nil,
      "success": false,
    })
	}

	booking.SetBooking(bookingInput)	
	
	result := databse.DB.Create(&booking)
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
		BookingID: int(booking.ID),
	}

  return c.JSON(fiber.Map{
    "message": "Successfully created the booking",
    "success": true,
    "data": bookingReturn,
  })
}



