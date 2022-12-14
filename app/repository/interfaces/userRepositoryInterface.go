package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
	"gitlab.com/cinco/app/response"
)

type UserRepositoryInterface interface {
	FindById(userUUID string) model.User
	Update(user model.User) error
	UserRegister(ctx *fiber.Ctx, params model.User) error
	GetUserByIdentity(ctx *fiber.Ctx, params string) (*model.User, error)
	SetRedis(key, val string, ttl int)
	DelRedis(key string)
	GetUserDetail(ctx *fiber.Ctx, user *response.ProfileDetail, params string) error
	CheckUser(ctx *fiber.Ctx, paramsUsername, paramsEmail string) (*model.User, error)
}
