package golib

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()

func BindNValidate(c echo.Context, out interface{}) error {
	err := c.Bind(out)
	if err != nil {
		return err
	}

	if err := validate.Struct(out); err != nil {
		er := map[string]string{}
		for _, v := range err.(validator.ValidationErrors) {
			field := ToSnakeCase(v.Field())
			msg := ""

			switch v.Tag() {
			case "required":
				msg = field + " is required"
			case "required_if":
				msg = field + " is required"
			case "email":
				msg = "not a valid email"
			case "len":
				msg = field + " " + "not in correct length"
			case "startswith":
				msg = field + " " + "not start with proper character"
			default:
				msg = field + " " + strings.ToLower(v.Tag())
			}

			er[field] = msg
		}
		b, err := json.Marshal(er)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", b)
	}
	return nil
}
