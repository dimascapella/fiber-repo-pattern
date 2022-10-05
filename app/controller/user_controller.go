package controller

import (
	"fiber-repo-pattern/app/model"
	"fiber-repo-pattern/app/request"
	"fiber-repo-pattern/app/service"
	"fiber-repo-pattern/helper"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Index(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type userController struct {
	userService service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		userService: service,
	}
}

func (handler *userController) Index(ctx *fiber.Ctx) error {
	var users []model.User = handler.userService.FindAll()
	result := helper.SuccessResponse(true, "List Data", users)
	return ctx.JSON(result)
}

func (handler *userController) Create(ctx *fiber.Ctx) error {
	var userCreateValidation request.UserCreateValidation
	err := ctx.BodyParser(&userCreateValidation)
	if err != nil {
		response := helper.ErrorResponse("Request Failed", err.Error(), helper.EmptyObj{})
		ctx.JSON(response)
		return err
	}
	user := handler.userService.Create(userCreateValidation)
	result := helper.SuccessResponse(true, "Data User Created", user)
	ctx.JSON(result)
	return nil
}

func (handler *userController) FindById(ctx *fiber.Ctx) error {
	id, _ := strconv.ParseUint(ctx.Params("id"), 10, 32)
	user := handler.userService.FindById(int(id))
	if (user == model.User{}) {
		res := helper.ErrorResponse("Data Not Found", "No Id Given", helper.EmptyObj{})
		ctx.JSON(res)
		return fiber.ErrBadRequest
	}
	result := helper.SuccessResponse(true, "Data User", user)
	ctx.JSON(result)
	return nil
}

func (handler *userController) Update(ctx *fiber.Ctx) error {
	id, _ := strconv.ParseUint(ctx.Params("id"), 10, 32)
	user := handler.userService.FindById(int(id))
	if (user == model.User{}) {
		res := helper.ErrorResponse("Data Not Found", "No Id Given", helper.EmptyObj{})
		ctx.JSON(res)
		return fiber.ErrBadRequest
	}
	var userUpdateValidation request.UserUpdateValidation
	userUpdateValidation.ID = ctx.Params("id")
	err := ctx.BodyParser(&userUpdateValidation)
	if err != nil {
		response := helper.ErrorResponse("Request Failed", err.Error(), helper.EmptyObj{})
		ctx.JSON(response)
		return err
	}
	updateUser := handler.userService.Update(userUpdateValidation)
	result := helper.SuccessResponse(true, "Data User Updated", updateUser)
	ctx.JSON(result)
	return nil
}

func (handler *userController) Delete(ctx *fiber.Ctx) error {
	id, _ := strconv.ParseUint(ctx.Params("id"), 10, 32)
	user := handler.userService.FindById(int(id))
	if (user == model.User{}) {
		res := helper.ErrorResponse("Data Not Found", "No Id Given", helper.EmptyObj{})
		ctx.JSON(res)
		return fiber.ErrBadRequest
	}
	deleteUser := handler.userService.Delete(user)
	result := helper.SuccessResponse(true, "Data User Deleted", deleteUser)
	ctx.JSON(result)
	return nil
}
