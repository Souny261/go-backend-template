package handler

import (
	"backend/internal/core/dto"
	"backend/internal/core/mappers"
	"backend/internal/core/ports/input"
	"backend/pkg/response"
	"backend/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type (
	UserHandler interface {

		// User
		GetUser(ctx *fiber.Ctx) error
		GetUserByID(ctx *fiber.Ctx) error
		UpdateUser(ctx *fiber.Ctx) error
		DeleteUser(ctx *fiber.Ctx) error
		UploadAvatar(ctx *fiber.Ctx) error
	}
	userHandler struct {
		userService   input.UserService
		mapperFactory *mappers.MapperFactory
	}
)

// UploadAvatar implements UserHandler.
func (h *userHandler) UploadAvatar(ctx *fiber.Ctx) error {
	userDTO := &dto.UserDTO{}
	if err := ctx.BodyParser(userDTO); err != nil {
		return response.NewErrorResponses(ctx, err)
	}
	user, err := h.mapperFactory.UserMapper().UserDTOToDomain(userDTO)
	if err != nil {
		return response.NewErrorResponses(ctx, err)
	}
	err = h.userService.UploadAvatar(ctx, user)
	if err != nil {
		return response.NewErrorResponses(ctx, err)
	}
	return response.NewSuccessMessagesResponse(ctx, "Avatar uploaded successfully")
}

// DeleteUser implements UserHandler.
func (h *userHandler) DeleteUser(ctx *fiber.Ctx) error {
	id, err := utils.StringToUint(ctx.Params("id"))
	if err != nil {
		return response.NewErrorResponses(ctx, err)
	}
	err = h.userService.DeleteUser(ctx.Context(), id)
	if err != nil {
		return response.NewErrorResponses(ctx, err)
	}
	return response.NewSuccessMessagesResponse(ctx, "User deleted successfully")
}

// UpdateUser implements UserHandler.
func (h *userHandler) UpdateUser(ctx *fiber.Ctx) error {
	userDTO := &dto.UserDTO{}
	if err := ctx.BodyParser(userDTO); err != nil {
		return response.NewErrorResponses(ctx, err)
	}
	user, err := h.mapperFactory.UserMapper().UserDTOToDomain(userDTO)
	if err != nil {
		return response.NewErrorResponses(ctx, err)
	}
	err = h.userService.UpdateUser(ctx.Context(), user)
	if err != nil {
		return response.NewErrorResponses(ctx, err)
	}
	return response.NewSuccessMessagesResponse(ctx, "User updated successfully")
}

// GetUserByID implements UserHandler.
func (h *userHandler) GetUserByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return response.NewErrorResponses(ctx, err)
	}
	user, err := h.userService.GetUserByID(ctx.Context(), uint(id))
	if err != nil {
		return response.NewErrorResponses(ctx, err)
	}
	res := h.mapperFactory.UserMapper().UserDomainToDTO(user)
	return response.NewSuccessResponse(ctx, res)
}

func (h *userHandler) GetUser(c *fiber.Ctx) error {
	users, err := h.userService.GetUsers(c.Context())
	if err != nil {
		return response.NewErrorResponses(c, err)
	}
	res := h.mapperFactory.UserMapper().DomainToDTOs(users)
	return response.NewSuccessResponse(c, res)
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(
	userService input.UserService,
) UserHandler {
	return &userHandler{
		userService:   userService,
		mapperFactory: mappers.NewMapperFactory(),
	}
}
