package controller

import (
	"net/http"
	"strconv"
	"time"

	"myapp-backend/controller/response"
	"myapp-backend/domain"
	"myapp-backend/service"

	"github.com/labstack/echo/v4"
)

type TransactionQueryController struct {
	queryService service.ITransactionQueryService
}

func NewTransactionQueryController(queryService service.ITransactionQueryService) *TransactionQueryController {
	return &TransactionQueryController{queryService: queryService}
}

func (c *TransactionQueryController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/transactions/history", c.GetHistory)
}

func (c *TransactionQueryController) GetHistory(ctx echo.Context) error {
	userIDStr := ctx.QueryParam("user_id")
	userID, _ := strconv.ParseInt(userIDStr, 10, 64)

	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	filter := domain.TransactionFilter{
		UserID: userID,
		Page:   page,
		Limit:  limit,
		SortBy: ctx.QueryParam("sort"),
	}

	if ctx.QueryParam("order") == "asc" {
		filter.Order = domain.SortAsc
	} else {
		filter.Order = domain.SortDesc
	}

	if from := ctx.QueryParam("from"); from != "" {
		t, _ := time.Parse(time.RFC3339, from)
		filter.FromTime = &t
	}

	if to := ctx.QueryParam("to"); to != "" {
		t, _ := time.Parse(time.RFC3339, to)
		filter.ToTime = &t
	}

	transactions, total, err := c.queryService.GetHistory(filter)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data":  transactions,
		"total": total,
		"page":  filter.Page,
		"limit": filter.Limit,
	})
}
