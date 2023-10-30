package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidatorMobile(filed validator.FieldLevel) bool {
	mobile := filed.Field().String()
	ok, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, mobile)
	if !ok {
		return false
	} else {
		return true
	}

}
