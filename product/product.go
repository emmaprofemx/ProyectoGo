package product

import (
	"github.com/devjaime/golangrest/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber"
)

type Productos struct {
	ID     uint   `gorm:"column:idProductos;primary_key;auto_increment"`
	Nombre string `json:"nombre"`
	Precio string `json:"precio"`
	Medida string `json:"medida"`
	Stock  int    `json:"stock"`
}

func GetProducts(c *fiber.Ctx) {
	db := database.DBConn
	var productos []Productos
	db.Find(&productos)
	c.JSON(productos)
}

func GetProduct(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var productos Productos
	db.Find(&productos, "idProductos = ?", id)
	c.JSON(productos)
}
func NewProduct(c *fiber.Ctx) {
	db := database.DBConn
	var productos Productos
	productos.Nombre = "Chocolate"
	productos.Precio = "50.00"
	productos.Medida = "Litro"
	productos.Stock = 10
	db.Create(&productos)
	c.JSON(productos)
}
func UpdateProduct(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var productos Productos
	db.First(&productos, id)

	if productos.ID == 0 {
		c.Status(fiber.StatusNotFound).Send("No se encontró el platillo")
		return
	}

	// Obtener los nuevos valores del cuerpo de la solicitud
	var newProducto Productos
	if err := c.BodyParser(&newProducto); err != nil {
		c.Status(fiber.StatusBadRequest).Send("Error al analizar la solicitud")
		return
	}

	// Actualizar los valores del platillo
	db.Model(&productos).Updates(&newProducto)

	c.JSON(productos)
}
