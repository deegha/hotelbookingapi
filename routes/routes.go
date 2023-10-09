package routes

import (
  "hotelbookingapi/controllers"
  "github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
  app.Post("/register", controllers.Register)
  app.Post("/login", controllers.Login)
  app.Post("/product", controllers.CreateProduct)
  app.Put("/product/:id", controllers.UpdateProduct)
  app.Post("/productSearch", controllers.ProductListHandler)
  app.Post("/booking", controllers.BookingCreateHandler)
  app.Put("/booking/:id", controllers.BookingUpdateHandler)
}
