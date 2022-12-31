package delivery

import (
	"net/http"
	"split_bills/helper"
	"split_bills/model"
	"split_bills/request"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

type billsDelivery struct {
	billUsecase model.BillsUseCase
}

type BillDelivery interface {
	Mount(group *echo.Group)
}

func NewBillDelivery(billUsecase model.BillsUseCase) BillDelivery {
	return &billsDelivery{billUsecase: billUsecase}
}

func (r *billsDelivery) Mount(group *echo.Group) {
	group.GET("", r.FetchBillHandler)
	group.POST("", r.StoreBillHandler)
	group.GET("/:id", r.DetailBillHandler)
	group.DELETE("/:id", r.DeleteBillHandler)
	group.PATCH("/:id", r.EditBillHandler)
}

func (r *billsDelivery) FetchBillHandler(c echo.Context) error {
	ctx := c.Request().Context()

	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	limitInt, _ := strconv.Atoi(limit)

	offsetInt, _ := strconv.Atoi(offset)

	billList, err := r.billUsecase.FetchBills(ctx, limitInt, offsetInt)

	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}

	return helper.ResponseSuccessJson(c, "success", billList)
}

func (r *billsDelivery) StoreBillHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.BillsRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error validation", errVal)
	}

	bill, err := r.billUsecase.StoreBills(ctx, &req)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return helper.ResponseSuccessJson(c, "success", bill)
}
func (r *billsDelivery) DetailBillHandler(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	IdInt, _ := strconv.Atoi(id)

	bill, err := r.billUsecase.GetByID(ctx, IdInt)

	if err != nil {
		return err
	}

	return helper.ResponseSuccessJson(c, "", bill)
}

func (r *billsDelivery) DeleteBillHandler(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	IdInt, _ := strconv.Atoi(id)

	err := r.billUsecase.DestroyBills(ctx, IdInt)

	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
	}

	return helper.ResponseSuccessJson(c, "", "")
}

func (r *billsDelivery) EditBillHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.BillsRequest

	if err := c.Bind(&req); err != nil {
		return err
	}

	err := c.Validate(req)
	if err != nil {
		return err
	}

	id := c.Param("id")

	IdInt, _ := strconv.Atoi(id)

	bill, err := r.billUsecase.EditBills(ctx, IdInt, &req)

	if err != nil {
		return err
	}

	return helper.ResponseSuccessJson(c, "Success Edit", bill)
}
