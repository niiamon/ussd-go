package ussd

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type core struct {
}

// MenuProcessor processes a menu rendered from Context.RenderMenu
func (cr *core) MenuProcessor(c *Context) *Response {
	str, err := c.DataBag.Get(coredataMenu)
	if err != nil {
		return c.Err(err)
	}
	menu := new(Menu)
	err = json.Unmarshal([]byte(str), menu)
	if err != nil {
		return c.Err(err)
	}
	errNotInMenu := errors.New(
		c.Request.Message + " is not in menu options.")
	choice, err := strconv.ParseInt(c.Request.Message, 10, 8)
	if err != nil {
		return c.Err(errNotInMenu)
	}
	if choice == 0 && menu.ZeroItem != nil {
		return c.Redirect(menu.ZeroItem.Route.Ctrl,
			menu.ZeroItem.Route.Action)
	}
	if int(choice) > len(menu.Items) {
		return c.Err(errNotInMenu)
	}
	item := menu.Items[choice-1]
	c.DataBag.Delete(coredataMenu)
	return c.Redirect(item.Route.Ctrl, item.Route.Action)
}

// FormInputDisplay displays inputs rendered from Context.RenderForm
func (cr *core) FormInputDisplay(c *Context) *Response {
	form, err := getForm(c)
	if err != nil {
		return c.Err(err)
	}
	input := form.Inputs[form.ProcessingPosition]
	displayName := StrEmpty
	if StrTrim(input.DisplayName) == StrEmpty {
		displayName = input.Name
	} else {
		displayName = input.DisplayName
	}
	msg := StrEmpty
	if StrTrim(form.Title) != StrEmpty {
		msg += form.Title + StrNewLine
	}
	if !input.HasOptions() {
		msg += fmt.Sprintf("Enter %v:"+StrNewLine, displayName)
	} else {
		msg += fmt.Sprintf("Select %v:"+StrNewLine, displayName)
		for i, option := range input.Options {
			value := StrEmpty
			if StrTrim(option.DisplayValue) == StrEmpty {
				value = option.Value
			} else {
				value = option.DisplayValue
			}
			msg += fmt.Sprintf("%d. %v"+StrNewLine, i+1, value)
		}
	}
	return c.Render(msg, "core", "FormInputProcessor")
}

func (cr *core) FormInputProcessor(c *Context) *Response {
	form, err := getForm(c)
	if err != nil {
		return c.Err(err)
	}
	input := form.Inputs[form.ProcessingPosition]
	key := input.Name
	value, err := getFormInputValue(c, input)
	if err != nil {
		return c.Err(err)
	}
	form.Data[key] = value
	if form.ProcessingPosition == len(form.Inputs)-1 {
		c.DataBag.Delete(coredataForm)
		c.FormData = form.Data
		return c.Redirect(form.Route.Ctrl, form.Route.Action)
	}
	form.ProcessingPosition++
	b, err := json.Marshal(form)
	if err != nil {
		return c.Err(err)
	}
	c.DataBag.Set(coredataForm, string(b))
	return c.Redirect("core", "FormInputDisplay")
}

func getForm(c *Context) (*Form, error) {
	str, err := c.DataBag.Get(coredataForm)
	if err != nil {
		return nil, err
	}
	form := new(Form)
	err = json.Unmarshal([]byte(str), form)
	if err != nil {
		return nil, err
	}
	return form, nil
}

func getFormInputValue(c *Context, input *Input) (string, error) {
	if !input.HasOptions() {
		return c.Request.Message, nil
	}
	errNotExist := fmt.Errorf(
		"Selected option %v does not exist.", c.Request.Message)
	choice, err := strconv.ParseInt(c.Request.Message, 10, 8)
	if err != nil {
		return StrEmpty, errNotExist
	}
	if int(choice) > len(input.Options) {
		return StrEmpty, errNotExist
	}
	return input.Options[choice-1].Value, nil
}
