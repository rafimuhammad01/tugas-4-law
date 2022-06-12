package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rafimuhammad01/api-gateway-service/entity"
	"github.com/rafimuhammad01/api-gateway-service/entity/dto"
	"github.com/rafimuhammad01/api-gateway-service/service"
	"github.com/rafimuhammad01/api-gateway-service/utils"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UpdateHandler struct {
	service service.UpdateService
}

func NewUpdateHandler(service service.UpdateService) *UpdateHandler {
	return &UpdateHandler{
		service: service,
	}
}

func (h *UpdateHandler) Create(ctx echo.Context) error {
	var request dto.CreateRequest
	err := ctx.Bind(&request)

	err = h.service.Create(request)
	if err != nil {
		if errors.Cause(err) == entity.ErrInvalidName || errors.Cause(err) == entity.ErrInvalidNPM {
			return ctx.JSON(http.StatusBadRequest, dto.APIResponse{
				Status: errors.Cause(err).Error(),
				Error:  utils.ErrorsParser(err),
			})
		}

		logrus.Error(err)
		return ctx.JSON(http.StatusInternalServerError, dto.APIResponse{
			Status: "internal server error",
		})
	}

	return ctx.JSON(http.StatusOK, dto.APIResponse{
		Status: "OK",
	})
}
