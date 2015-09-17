package validator

import (
	"errors"
	"fmt"

	"github.com/samora/ussd-go/Godeps/_workspace/src/github.com/asaskevich/govalidator"
)

// Map contains all validators
var Map = map[string]Validator{
	"length":  Length,
	"numeric": Numeric,
	"integer": Integer,
	"float":   Float,
	"range":   Range,
}

// Validator is the function signature for validators
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
		if min == max {
			return fmt.Errorf("%s must have length of %d", name, max)
		}
		return fmt.Errorf("%s must have length ranging from %d to %d", name, min, max)
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
		return fmt.Errorf("%s must only contain numbers", name)
	}
	return nil
}

// Float verifies value is a float/decimal number
func Float(name, value string, args ...string) error {
	if ok := govalidator.IsFloat(value); !ok {
		return fmt.Errorf("%s must be a float/decimal", name)
	}
	return nil
}

// Range verifies value falls in range provided
func Range(name, value string, args ...string) error {
	if len(args) != 2 {
		panic(errors.New("Min and max must be specified"))
	}
	v, err := govalidator.ToFloat(value)
	if err != nil {
		return err
	}
	min, err := govalidator.ToFloat(args[0])
	if err != nil {
		return err
	}
	max, err := govalidator.ToFloat(args[1])
	if err != nil {
		return err
	}
	if ok := govalidator.InRange(v, min, max); !ok {
		return fmt.Errorf("%s must range from %f to %f", name, min, max)
	}
	return nil
}
