package pkg

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var ElectricityIdValidator validator.Func = func(fieldLevel validator.FieldLevel) bool {
	electricityId := fieldLevel.Field().String()

	validElectricityId := regexp.MustCompile("^[0-9]{11}$")
	return electricityId != "" && validElectricityId.MatchString(electricityId)
}
