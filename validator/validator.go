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
type Validator func(string, ...string) error

// Length verifies min <= value <= max.
// Where min and max are first and second args respectively.
// max is optional
func Length(value string, args ...string) error {
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
			return fmt.Errorf("Must have min length of %d", min)
		}
		if min == max {
			return fmt.Errorf("Must have length of %d", max)
		}
		return fmt.Errorf("Must have length ranging from %d to %d", min, max)
	}
	return nil
}

// Integer verifies value is a valid integer
func Integer(value string, args ...string) error {
	if ok := govalidator.IsInt(value); !ok {
		return errors.New("Must be an integer")
	}
	return nil
}

// Numeric verifies value only contains chars +,-,0-9
func Numeric(value string, args ...string) error {
	if ok := govalidator.IsNumeric(value); !ok {
		return errors.New("Must only contain numbers")
	}
	return nil
}

// Float verifies value is a float/decimal number
func Float(value string, args ...string) error {
	if ok := govalidator.IsFloat(value); !ok {
		return errors.New("Must be a float/decimal")
	}
	return nil
}

// Range verifies value falls in range provided
func Range(value string, args ...string) error {
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
		return fmt.Errorf("Must range from %.2f to %.2f", min, max)
	}
	return nil
}
