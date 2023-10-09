
package databse

import (
	"hotelbookingapi/models"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(){
	connection, error := gorm.Open(mysql.Open("root:root@/hotels?parseTime=true"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if error != nil {
		panic("couldn't connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{}, &models.Product{}, &models.Amenities{}, &models.Media{}, &models.ProductMedia{}, &models.Booking{})
}
