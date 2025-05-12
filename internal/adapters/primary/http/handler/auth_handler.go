package handler

import (
	"backend/internal/core/dto"
	"backend/internal/core/mappers"
	"backend/internal/core/ports/input"
	"backend/pkg/jwt"
	"backend/pkg/response"
	"backend/pkg/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type (
	AuthHandler interface {
		Login(ctx *fiber.Ctx) error
		WhoAmI(ctx *fiber.Ctx) error
		RefreshToken(ctx *fiber.Ctx) error
	}
	authHandler struct {
		authService   input.AuthService
		mapperFactory *mappers.MapperFactory
	}
)

// RefreshToken implements AuthHandler.
func (a *authHandler) RefreshToken(ctx *fiber.Ctx) error {
	jwt, err := jwt.GenerateRefreshToken(ctx)
	if err != nil {
		return response.NewErrorResponses(ctx, err)
	}
	data := a.mapperFactory.UserMapper().TokenPairToJWTDTO(jwt)
	return response.NewSuccessResponse(ctx, data)
}

// WhoAmI implements AuthHandler.
func (a *authHandler) WhoAmI(ctx *fiber.Ctx) error {
	tokenDecoded, err := jwt.GetOwnerAccessToken(ctx)
	if err != nil {
		return response.NewErrorResponses(ctx, err)
	}
	user, err := a.authService.WhoAmI(ctx.Context(), uint(tokenDecoded.ID))
	if err != nil {
		return response.NewErrorResponses(ctx, err)
	}
	return response.NewSuccessResponse(ctx, a.mapperFactory.UserMapper().UserDomainToDTO(user))
}

// Login implements AuthHandler.
func (a *authHandler) Login(ctx *fiber.Ctx) error {
	var userDTO dto.UserDTO
	if err := ctx.BodyParser(&userDTO); err != nil {
		return response.NewErrorResponses(ctx, err)
	}
	user, err := a.authService.Login(ctx.Context(), userDTO.Username)
	if err != nil {
		if err.Error() == "data not found" {
			return response.NewErrorResponses(ctx, errors.New("username or password is incorrect"))
		}
		return response.NewErrorResponses(ctx, err)
	}
	if !user.Status {
		return response.NewErrorResponses(ctx, errors.New("user is inactive"))
	}
	err = utils.VerifyPassword(string(user.PasswordHash), userDTO.Password)
	if err != nil {
		return response.NewErrorResponses(ctx, errors.New("username or password is incorrect"))
	}
	jwt, err := jwt.GenerateJWTToken(int(user.ID))
	if err != nil {
		return response.NewErrorResponses(ctx, err)
	}
	a.authService.LastUserActive(ctx.Context(), user.ID)
	data := a.mapperFactory.UserMapper().TokenPairToJWTDTO(jwt)
	return response.NewSuccessResponse(ctx, data)
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(
	authService input.AuthService,
) AuthHandler {
	return &authHandler{
		authService:   authService,
		mapperFactory: mappers.NewMapperFactory(),
	}
}
