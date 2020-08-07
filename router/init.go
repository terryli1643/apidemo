package router

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

const (
	PROFILE = "profile"
	HANDLER = "Handler"
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
		log.WithFields(log.Fields{
			"FieldError.ActualTag":       fieldError.ActualTag(),
			"FieldError.Field":           fieldError.Field(),
			"FieldError.StructNamespace": fieldError.StructNamespace(),
			"FieldError.Kind":            fieldError.Kind(),
			"FieldError.Namespace":       fieldError.Namespace(),
			"FieldError.StructField":     fieldError.StructField(),
			"FieldError.Param":           fieldError.Param(),
			"FieldError.Tag":             fieldError.Tag(),
			"FieldError.Type":            fieldError.Type(),
			"FieldError.Value":           fieldError.Value(),
			"FieldError.Translate":       fieldError.Translate,
		}).Debug("ValidationErrors")
		if fieldError != nil {
			return fieldError.Field() + ":" + combineErrorKey(fieldError)
		}
	}
	return ""
}

func combineErrorKey(fieldError validator.FieldError) string {
	var message, namespace, actualTag string
	namespace = fieldError.Namespace()
	actualTag = fieldError.ActualTag()

	switch actualTag {
	case "required":
		message = "不能为空"
		// case "max":
		// 	message = "最大不能超过" + fieldError.Param()
		// case "min":
		// 	message = "不能小于")
		// case "lt":
		// 	message = alert.T("key_alert_fielderror_field_lessthan", map[string]interface{}{
		// 		"nu": fieldError.Param(),
		// 	})
		// case "gt":
		// 	message = alert.T("key_alert_fielderror_field_greaterthan", map[string]interface{}{
		// 		"nu": fieldError.Param(),
		// 	})
		// case "email":
		// 	message = alert.T("key_alert_fielderror_field_email")
		// case "alphanum":
		// 	message = alert.T("key_alert_fielderror_field_alphanum")
		// case "numeric":
		// 	message = alert.T("key_alert_fielderror_field_numeric")
		// case "url":
		// 	message = alert.T("key_alert_fielderror_field_url")
		// case "pwdmatch":
		// 	message = alert.T("key_alert_fielderror_field_pwdmatch")
		// case "timerange":
		// 	message = alert.T("key_alert_fielderror_field_timerange")
		// case "moneyrange":
		// 	message = alert.T("key_alert_fielderror_field_moneyrange")
		// case "ip":
		// 	message = alert.T("key_alert_fielderror_field_ip")
		// case "ipv4":
		// 	message = alert.T("key_alert_fielderror_field_ip")
		// case "ipv6":
		// 	message = alert.T("key_alert_fielderror_field_ip")
		// default:
		// 	Log.WithField("acturelTag", actualTag).Error("illegal acturelTag")
	}
	log.WithFields(log.Fields{
		"namespace": namespace,
		"actualTag": actualTag,
		"message":   message,
	}).Debug("TAlert")

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
