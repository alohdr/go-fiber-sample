package interfaces

import (
	"gitlab.com/cinco/app/response"
	"time"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
)

type CashflowRepositoryInterface interface {
	PostTransaction(ctx *fiber.Ctx, body *model.Cashflow) error
	FindByAccount(userUUID string, tipe string, startDate time.Time, endDate time.Time) ([]model.Cashflow, error)
	FindTotal(userUUID string, startDate time.Time, endDate time.Time) (response.Total, error)
	DeleteCashflow(ctx *fiber.Ctx, params string) error
	RepoEditCashFlow(ctx *fiber.Ctx, editcashflow *model.Cashflow, params string) error
	RepoUpdateBalance(ctx *fiber.Ctx, updatebalance int, paramsIdAccount string) error
	GetHistoryandAmountBefore(ctx *fiber.Ctx, params string) (int, string, error)
}
