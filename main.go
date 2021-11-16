package main

import (
	"log"
	_ "github.com/MAHakeem-MK/go_server/docs"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
    gorm.Model
    Code  string 
    Price string 
}
// GetProduct godoc
// @Summary Show a product
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Failure 404 {object} string
// @Success 200 {object} Product
// @Router /product/{id} [get]
func GetProduct(c *fiber.Ctx) error {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    var product Product
    el := db.First(&product, c.Params("id"))
    if (el != nil) {
        return c.SendStatus(404)
    }
    return c.JSON(product)
}

func GetProducts(c *fiber.Ctx) error {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    var products Product
    db.Find(&products)
    return c.JSON(products)
}

// @title Fiber Example API
// @version 1.0
// @host localhost:3000
// @BasePath /
func main() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&Product{})
    app := fiber.New()

    app.Get("/docs/*", swagger.Handler)

    app.Post("/product/:id/price/:price",func (c *fiber.Ctx) error {
        db.Create(&Product{Code: c.Params("id"), Price: c.Params("price")})
        return c.SendString("Product added")
    })
    
    app.Get("/product/:id", GetProduct)

    app.Get("/products",GetProducts)

    log.Fatal(app.Listen(":3000"))
}