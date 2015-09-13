package ussd

// Form is a USSD form.
type Form struct {
	Title              string
	Route              route
	ProcessingPosition int
	Data               map[string]string
	Inputs             []*Input
}

// NewForm creates a new form.
func NewForm(title string) *Form {
	return &Form{
		Title:  title,
		Data:   make(map[string]string),
		Inputs: make([]*Input, 0),
	}
}

// AddInput adds an input to USSD form.
func (f *Form) AddInput(name, displayName string,
	options ...*Option) *Form {
	input := NewInput(name, displayName)
	for _, option := range options {
		input.Options = append(input.Options, option)
	}
	f.Inputs = append(f.Inputs, input)
	return f
}

// Input for USSD form.
type Input struct {
	Name, DisplayName string
	Options           []*Option
}

// NewInput creates new form input.
func NewInput(name, displayName string) *Input {
	return &Input{
		Name:        name,
		DisplayName: displayName,
		Options:     make([]*Option, 0),
	}
}

// HasOptions checks if input has options.
func (i Input) HasOptions() bool {
	return len(i.Options) > 0
}

// Option for USSD form select field.
type Option struct {
	Value, DisplayValue string
}

// NewOption create a USSD input option.
func NewOption(value, displayValue string) *Option {
	return &Option{
		Value: value, DisplayValue: displayValue,
	}
}
