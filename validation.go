package golib

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func BindNValidate(c echo.Context, out interface{}) error {
	err := c.Bind(out)
	if err != nil {
		return err
	}

	if err != nil {
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

			return errors.New(msg)
		}
		return nil
	}
	return nil
}
