package main

import ( 
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
  "hotelbookingapi/routes"
  "hotelbookingapi/database"
)

func main() {
  app:= fiber.New()
	databse.Connect()
  app.Use(cors.New(cors.Config{
    AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
    AllowOrigins:     "*",
    AllowCredentials: true,
    AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
  }))
  routes.Setup(app)
  app.Listen(":8010")
}
