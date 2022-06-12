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
	"strconv"
)

type ReadHandler struct {
	service service.ReadService
}

func NewReadHandler(service service.ReadService) *ReadHandler {
	return &ReadHandler{
		service: service,
	}
}

func (h *ReadHandler) ReadWithTransaction(ctx echo.Context) error {
	var studentInfo dto.ReadRequest
	studentInfo.NPM = ctx.Param("npm")

	trxIDParam := ctx.Param("trxID")
	trxID, err := strconv.Atoi(trxIDParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.APIResponse{
			Status: entity.ErrInvalidTrxID.Error(),
			Error:  "transaction id must be positive integer",
		})
	}
	studentInfo.TransactionID = trxID

	resp, err := h.service.ReadWithTransaction(studentInfo)
	if err != nil {
		if errors.Cause(err) == entity.ErrStudentNotFound {
			return ctx.JSON(http.StatusNotFound, dto.APIResponse{
				Status: errors.Cause(err).Error(),
				Error:  utils.ErrorsParser(err),
			})
		}

		if errors.Cause(err) == entity.ErrInvalidTrxID {
			return ctx.JSON(http.StatusBadRequest, dto.APIResponse{
				Status: errors.Cause(err).Error(),
				Error:  utils.ErrorsParser(err),
			})
		}

		logrus.Info(err)
		return ctx.JSON(http.StatusInternalServerError, dto.APIResponse{
			Status: "internal server error",
		})
	}

	return ctx.JSON(http.StatusOK, dto.APIResponse{
		Status: "OK",
		Data:   resp,
	})
}

func (h *ReadHandler) Read(ctx echo.Context) error {
	var studentInfo dto.ReadRequest
	studentInfo.NPM = ctx.Param("npm")

	resp, err := h.service.Read(studentInfo)
	if err != nil {
		if errors.Cause(err) == entity.ErrStudentNotFound {
			return ctx.JSON(http.StatusNotFound, dto.APIResponse{
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
		Data:   resp,
	})
}
