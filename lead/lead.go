package lead

import (
	"github.com/Patkeenan/go-crm-project/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
	return nil
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON((lead))
	return nil

}

func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		return c.Status(503).Send([]byte(err.Error()))
	}
	db.Create(&lead)
	c.JSON(lead)
	return nil
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		return c.Status(500).Send([]byte("No lead with that ID"))

	}
	db.Delete(&lead)
	c.Send([]byte("Lead successfully deleted"))
	return nil
}
