package config

import (
	"fiber-repo-pattern/app/controller"
	"fiber-repo-pattern/app/repository"
	"fiber-repo-pattern/app/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = SetupConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	userService    service.UserService       = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService)
)

func Initalize(router *fiber.App) {
	users := router.Group("/users")
	users.Get("/", userController.Index)
	users.Post("/", userController.Create)
	users.Get("/:id", userController.FindById)
	users.Put("/:id", userController.Update)
	users.Delete("/:id", userController.Delete)

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})
}
