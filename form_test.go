package ussd

import (
	"testing"

	"github.com/samora/ussd-go/Godeps/_workspace/src/github.com/stretchr/testify/suite"
)

type FormSuite struct {
	suite.Suite
	form *Form
}

func (f *FormSuite) SetupSuite() {
	form := NewForm("User Registration")
	form.Input("Name", StrEmpty)
	form.Input("Sex", StrEmpty,
		form.Option("M", "Male"),
		form.Option("F", "Female"))
	form.Input("Age", StrEmpty).Validate("integer")
	f.form = form
}

func (f FormSuite) TestForm() {
	form := f.form

	f.Equal(len(form.Inputs), 3)
	sex := form.Inputs[1]
	age := form.Inputs[2]
	f.Equal(len(sex.Options), 2)
	f.Equal(len(age.Validators), 1)
}

func TestFormSuite(t *testing.T) {
	suite.Run(t, new(FormSuite))
}
