package ussd

import "github.com/samora/ussd-go/validator"

// Form is a USSD form.
type Form struct {
	Title, ValidationMessage string
	Route                    route
	ProcessingPosition       int
	Data                     FormData
	Inputs                   []input
}

// NewForm creates a new form.
func NewForm(title string) *Form {
	return &Form{
		Title:  StrTrim(title),
		Data:   make(FormData),
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

// Validate input. See validator.Map for available validators.
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

type input struct {
	Name, DisplayName string
	Options           []option
	Validators        []validatorData
}

type option struct {
	Value, DisplayValue string
}

type validatorData struct {
	Key  string
	Args []string
}

func newInput(name, displayName string) input {
	return input{
		Name:        name,
		DisplayName: displayName,
		Options:     make([]option, 0),
		Validators:  make([]validatorData, 0),
	}
}

func (i input) hasOptions() bool {
	return len(i.Options) > 0
}
