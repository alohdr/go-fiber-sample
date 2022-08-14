package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/app/model"
	"time"
)

type CashflowServiceInterface interface {
	AddTransaction(ctx *fiber.Ctx, body model.Cashflow) error
	FindTransactionLog(userUUID string, tipe string, startDate time.Time, endDate time.Time) ([]model.Cashflow, error)
	EditCashflow(ctx *fiber.Ctx, body *model.Cashflow, reqUpdate *model.Account, params, paramsIdAccount string) (*model.Cashflow, error)
	DeleteCashflow(ctx *fiber.Ctx, cashflowid string, paramsIdAccount string) (*model.Cashflow, error)
}
