package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/terryli1643/apidemo/libs/security"
)

type (
	WrongGoogleToken struct {
		Err error
	}
	WrongCaptchaCode struct {
		Err error
	}
	MissGoogleToken struct {
		Err error
	}
	VisitLimit struct {
		Err error
	}

	BalanceFrozenErr struct {
		Err error
	}
)

func (e WrongGoogleToken) Error() string {
	return e.Err.Error()
}
func (e WrongCaptchaCode) Error() string {
	return e.Err.Error()
}
func (e MissGoogleToken) Error() string {
	return e.Err.Error()
}
func (e VisitLimit) Error() string {
	return e.Err.Error()
}
func (e BalanceFrozenErr) Error() string {
	return e.Err.Error()
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				switch err.Err.(type) {
				case security.WrongUserNamePasswordError:
					newGenError(c, "用户名密码错误")
				case security.NotPromissionError:
					newGenError(c, "权限不足")
				case security.AccountLockedError:
					newGenError(c, "账号已锁定")
				case security.AccountExpiredError:
					newGenError(c, "账号已过期")
				case WrongGoogleToken:
					new400Error(c, "key_alert_error_token_error")
				case WrongCaptchaCode:
					newGenError(c, "key_alert_error_captcha_error")
				case MissGoogleToken:
					newGenError(c, "key_alert_error_captcha_error")
				case VisitLimit:
					newGenError(c, "key_alert_ip_not_allowed_error")
				case BalanceFrozenErr:
					newForbiddenError(c, "key_alert_error_balancefrozen")
				}
				break
			}
		}
	}
}

func newGenError(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, GenericMessageBody{
		Message: message,
	})
	c.Abort()
}

func new400Error(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, GenericMessageBody{
		Message: message,
	})
	c.Abort()
}

func newForbiddenError(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, GenericMessageBody{
		Message: message,
	})
	c.Abort()
}

type GenericSuccess struct {
	//in: body
	Body GenericMessageBody
}

// swagger:model
type GenericMessageBody struct {
	Message string
}
