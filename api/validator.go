package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/syafrin-ibrahim/mybank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		//check curreny is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
