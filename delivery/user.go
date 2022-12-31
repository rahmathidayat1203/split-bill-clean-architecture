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

type userDelivery struct {
	userUsercase model.UserUsecase
}

type UserDelivery interface {
	Mount(group *echo.Group)
}

func NewUserDelivery(userUsecase model.UserUsecase) UserDelivery {
	return &userDelivery{userUsercase: userUsecase}
}

func (r *userDelivery) Mount(group *echo.Group) {
	group.GET("", r.FetchUserHandler)
	group.POST("", r.StoreUserHandler)
	group.GET("", r.DetailUserHandler)
	group.DELETE("", r.DeleteUserHandler)
	group.PATCH("", r.EditUserHandler)
	group.POST("", r.BillsTotalHandler)
}
func (r *userDelivery) FetchUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	limitConvInt, _ := strconv.Atoi(limit)
	offsetConvInt, _ := strconv.Atoi(offset)

	userList, err := r.userUsercase.FetchUser(ctx, limitConvInt, offsetConvInt)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return helper.ResponseSuccessJson(c, "success", userList)

}

func (r *userDelivery) StoreUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.UserRequest

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error validation", errVal)
	}

	user, err := r.userUsercase.StoreUser(ctx, &req)

	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}
	return helper.ResponseSuccessJson(c, "success", user)

}
func (r *userDelivery) DetailUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	IdInt, _ := strconv.Atoi(id)

	user, err := r.userUsercase.GetByID(ctx, IdInt)

	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}

	return helper.ResponseSuccessJson(c, "", user)
}
func (r *userDelivery) DeleteUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	IdInt, _ := strconv.Atoi(id)

	err := r.userUsercase.DestroyUser(ctx, IdInt)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
	}

	return helper.ResponseSuccessJson(c, "", "")
}
func (r *userDelivery) EditUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.UserRequest

	if err := c.Bind(&req); err != nil {
		return helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())
	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error validation", errVal)
	}

	id := c.Param("id")
	IdInt, _ := strconv.Atoi(id)

	user, err := r.userUsercase.EditUser(ctx, IdInt, &req)

	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
	}
	return helper.ResponseSuccessJson(c, "Success edit", user)
}
func (r *userDelivery) BillsTotalHandler(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.BillsTotalRequest
	if err := c.Bind(&req); err != nil {
		return helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())
	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error validation", errVal)
	}
	err := r.userUsercase.BillsTotal(ctx, &req)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
	}

	return helper.ResponseSuccessJson(c, "Success Process Bills", "")

}
