package validator

import (
	"errors"
	"fmt"

	"github.com/samora/ussd-go/Godeps/_workspace/src/github.com/asaskevich/govalidator"
)

type Validator func(string, string, ...string) error

// Length verifies min <= value <= max.
// Where min and max are first and second args respectively.
// max is optional
func Length(name, value string, args ...string) error {
	if len(args) == 0 {
		panic(errors.New("Invalid args. Min length must be specified"))
	}
	min, err := govalidator.ToInt(args[0])
	if err != nil {
		return err
	}
	var max int64
	if len(args) == 2 {
		max, err = govalidator.ToInt(args[1])
		if err != nil {
			return err
		}
	}
	length := int64(len(value))
	if length < min || (max != 0 && length > max) {
		if max == 0 {
			return fmt.Errorf("%s must have min length of %d", name, min)
		}
		return fmt.Errorf("%s must have length %d to %d", name, min, max)
	}
	return nil
}

// Integer verifies value is a valid integer
func Integer(name, value string, args ...string) error {
	if ok := govalidator.IsInt(value); !ok {
		return fmt.Errorf("%s must be an integer", name)
	}
	return nil
}

// Numeric verifies value only contains chars +,-,0-9
func Numeric(name, value string, args ...string) error {
	if ok := govalidator.IsNumeric(value); !ok {
		return fmt.Errorf("%s can only contain numbers", name)
	}
	return nil
}

// Float verifies value is a float/decimal number
func Float(name, value string, args ...string) error {
	if ok := govalidator.IsFloat(value); !ok {
		return fmt.Errorf("%s must be a float/decimal number")
	}
	return nil
}
