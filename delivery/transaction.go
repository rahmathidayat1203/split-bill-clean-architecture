package delivery

import (
	"split_bills/helper"
	"split_bills/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

type transactionDelivery struct {
	transactionUsecase model.TransactionsUsecase
}

type TransactionDelivery interface {
	Mount(group *echo.Group)
}

func NewTransactionDelvery(transactionUsecase model.TransactionsUsecase) TransactionDelivery {
	return &transactionDelivery{transactionUsecase: transactionUsecase}
}

func (r *transactionDelivery) Mount(group *echo.Group) {
	group.GET("", r.FetchTransactionHandler)
}

func (r *transactionDelivery) FetchTransactionHandler(c echo.Context) error {
	ctx := c.Request().Context()

	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	transaction, i, err := r.transactionUsecase.Fetch(ctx, limitInt, offsetInt)

	if err != nil {
		return helper.ResponseErrorJson(c, i, err)
	}

	return helper.ResponseSuccessJson(c, "success", transaction)
}
