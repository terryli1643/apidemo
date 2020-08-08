package router

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
	"github.com/terryli1643/apidemo/libs/logger"
)

const (
	PROFILE = "profile"
	HANDLER = "Handler"
)

var (
	log = logger.InitLog()
)

type CommonResp struct {
	Code    int         `json:"code"`
	Body    interface{} `json:"body,omitempty"`
	Message string      `json:"message"`
}

func bindID(c *gin.Context) (int64, error) {
	var idstr string
	idstr, ok := c.GetQuery("ID")
	if !ok {
		return 0, errors.New("ID is required")
	}
	id64, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return id64, nil
}

func newClientError(c *gin.Context, err error) {
	log.Error(err)
	if e, ok := err.(validator.ValidationErrors); ok {
		c.JSON(http.StatusBadRequest, CommonResp{
			Code:    http.StatusBadRequest,
			Message: parseFieldError(e),
		})
		return
	}
	c.JSON(http.StatusBadRequest, CommonResp{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	})
}

func parseFieldError(validationErrors validator.ValidationErrors) string {
	for _, fieldError := range validationErrors {
		if fieldError != nil {
			return fieldError.Field() + ":" + combineErrorKey(fieldError)
		}
	}
	return ""
}

func combineErrorKey(fieldError validator.FieldError) string {
	var message, actualTag string
	actualTag = fieldError.ActualTag()

	switch actualTag {
	case "required":
		message = "不能为空"
	default:
	}
	return message
}

func newServerError(c *gin.Context, err error) {
	log.Error(err)
	c.JSON(http.StatusInternalServerError, CommonResp{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	})
}

func newSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, CommonResp{
		Code:    http.StatusOK,
		Message: "success",
	})
}

func PercentageForPrint(d decimal.Decimal) decimal.Decimal {
	return d.Mul(decimal.NewFromFloat(100))
}

func PercentageForSave(d decimal.Decimal) decimal.Decimal {
	if d != decimal.Zero {
		return d.Div(decimal.NewFromFloat(100))
	}
	return decimal.Zero
}
