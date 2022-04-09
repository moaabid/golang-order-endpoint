package routes

import (
	"fmt"

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

func GetUsers(c *fiber.Ctx) error {

	users := []model.User{}

	database.Database.DB.Find(&users)

	responseUsers := []UserSerializer{}

	for _, user := range users {
		responseUser := CreateReponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)

}
func GetUser(c *fiber.Ctx) error {

	var user model.User

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid user id")
	}

	database.Database.DB.Find(&user, "id=?", id)

	if user.Id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON("User Does not exists")
	}

	fmt.Println(user)

	responseUser := CreateReponseUser(user)

	return c.Status(200).JSON(responseUser)

}

func UpdateUser(c *fiber.Ctx) error {

	var user model.User

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid user id")
	}

	database.Database.DB.Find(&user, "id=?", id)

	if user.Id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON("User Does not exists")
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updateUser UpdateUser

	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateUser.FirstName
	user.LastName = updateUser.LastName

	database.Database.DB.Save(&user)

	responseUser := CreateReponseUser(user)

	return c.Status(200).JSON(responseUser)

}

func DeleteUser(c *fiber.Ctx) error {

	var user model.User

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid user id")
	}
	database.Database.DB.Find(&user, "id=?", id)

	if user.Id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON("User Does not exists")
	}
	if err := database.Database.DB.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("User deleted successfully")

}
