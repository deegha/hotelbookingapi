package controllers

import (
	"fmt"
	databse "hotelbookingapi/database"
	"hotelbookingapi/models"
	"hotelbookingapi/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SearchProducts(db *gorm.DB, searchQuery models.ProductSearch) ([]models.Product, error) {
    var products []models.Product

	  query := db.Model(&models.Product{}).Preload("Media")
	 
	 if searchQuery.SearchQuery != "" {
        query = query.Where("title LIKE ?", "%"+searchQuery.SearchQuery+"%")
    }

    if searchQuery.Limit > 0 {
        query = query.Limit(searchQuery.Limit)
    }

    if searchQuery.Offset > 0 {
        query = query.Offset(searchQuery.Offset)
    }

    if searchQuery.Price > 0 {
        query = query.Where("price <= ?", searchQuery.Price)
    }

    if searchQuery.RatingScore > 0 {
        query = query.Where("rating_score >= ?", searchQuery.RatingScore)
    }

	   // Adding conditions based on check-in and check-out times
    if !searchQuery.CheckInTime.IsZero() {
        query = query.Where("id NOT IN (SELECT product_id FROM bookings WHERE ? < check_out_time AND ? > check_in_time)", searchQuery.CheckOutTime, searchQuery.CheckInTime)
    }
	 if err := query.Find(&products).Error; err != nil {
        return nil, err
    }
   
    return products, nil
}

func ProductListHandler(c *fiber.Ctx) error {
	
	var productSearch models.ProductSearch 
	if err := c.BodyParser(&productSearch); err != nil {
    c.Status(fiber.StatusBadRequest)
    return c.JSON(fiber.Map{
      "message": "Could not parse the data",
      "data": nil,
      "success": false,
    })
  }
	
	products, error := SearchProducts(databse.DB, productSearch)
	
	for i := 0; i < len(products); i++ {
		nights, _ := utils.CalculateNumberOfNights(productSearch.CheckInTime, productSearch.CheckOutTime)
		fmt.Print(nights, "nights")
		products[i].Price = products[i].Price * nights 
	}

	if error != nil {
		return c.JSON(fiber.Map{
			"message": "Couldn't fetch the records",
			"data": nil,
			"success": false,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Succuessful",
		"data": products,
		"success": true,
	})
}

