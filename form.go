package ussd

import "github.com/samora/ussd-go/validator"

// Form is a USSD form.
type Form struct {
	Title, ValidationMessage string
	Route                    route
	ProcessingPosition       int
	Data                     map[string]string
	Inputs                   []input
}

// NewForm creates a new form.
func NewForm(title string) *Form {
	return &Form{
		Title:  StrTrim(title),
		Data:   make(map[string]string),
		Inputs: make([]input, 0),
	}
}

// Input adds an input to USSD form.
func (f *Form) Input(name, displayName string,
	options ...option) *Form {
	input := newInput(StrTrim(name), StrTrim(displayName))
	for _, option := range options {
		input.Options = append(input.Options, option)
	}
	f.Inputs = append(f.Inputs, input)
	return f
}

func (f *Form) Validate(validatorKey string, args ...string) *Form {
	validatorKey = StrTrim(StrLower(validatorKey))
	if _, ok := validator.Map[validatorKey]; !ok {
		panic(&validatorDoesNotExistError{validatorKey})
	}
	i := len(f.Inputs) - 1
	input := f.Inputs[i]
	input.Validators = append(input.Validators, validatorData{
		Key:  validatorKey,
		Args: args,
	})
	f.Inputs[i] = input
	return f
}

// Option creates a USSD input option.
func (f Form) Option(value, displayValue string) option {
	return option{
		Value: StrTrim(value), DisplayValue: StrTrim(displayValue),
	}
}

// input for USSD form.
type input struct {
	Name, DisplayName string
	Options           []option
	Validators        []validatorData
}

// option for input
type option struct {
	Value, DisplayValue string
}

type validatorData struct {
	Key  string
	Args []string
}

// newInput creates new form input.
func newInput(name, displayName string) input {
	return input{
		Name:        name,
		DisplayName: displayName,
		Options:     make([]option, 0),
		Validators:  make([]validatorData, 0),
	}
}

// hasOptions checks if input has options.
func (i input) hasOptions() bool {
	return len(i.Options) > 0
}
