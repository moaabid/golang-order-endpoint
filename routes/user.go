package routes

import (
	"githu.com/moaabid/golang-order-endpoint/database"
	"githu.com/moaabid/golang-order-endpoint/model"
	"github.com/gofiber/fiber/v2"
)

type UserSerializer struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateReponseUser(user model.User) UserSerializer {
	return UserSerializer{Id: user.Id, FirstName: user.FirstName, LastName: user.LastName}
}

func CreateUser(c *fiber.Ctx) error {

	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.DB.Create(&user)

	responseUser := CreateReponseUser(user)

	return c.Status(200).JSON(responseUser)

}
