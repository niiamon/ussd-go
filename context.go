package ussd

import (
	"encoding/json"
)

// Context of current USSD session.
type Context struct {
	Request  *Request
	DataBag  *DataBag
	Data     Data
	FormData FormData
}

type FormData map[string]string

func (f FormData) Exists() bool {
	return len(f) > 0
}

// Release USSD session.
func (c *Context) Release(msg string) *Response {
	r := new(Response)
	r.Message = StrTrim(msg)
	r.Release = true
	return r
}

// Render to USSD client. Expects USSD client to respond with
// input which will be mapped to specified route.
func (c *Context) Render(msg, ctrl, action string) *Response {
	r := new(Response)
	r.Message = StrTrim(msg)
	r.route = route{StrTrim(ctrl), StrTrim(action)}
	return r
}

// RenderMenu displays a menu and does appropriate routing to
// user's response.
func (c *Context) RenderMenu(menu *Menu) *Response {
	b, err := json.Marshal(menu)
	if err != nil {
		return c.Err(err)
	}
	c.DataBag.Set(coredataMenu, string(b))
	return c.Render(menu.Render(), "core", "MenuProcessor")
}

// RenderForm starts the form input collection process.
func (c *Context) RenderForm(form *Form, ctrl, action string) *Response {
	form.Route = route{StrTrim(ctrl), StrTrim(action)}
	b, err := json.Marshal(form)
	if err != nil {
		return c.Err(err)
	}
	c.DataBag.Set(coredataForm, string(b))
	return c.Redirect("core", "FormInputDisplay")
}

// Redirect to another route to continue processing request.
func (c *Context) Redirect(ctrl, action string) *Response {
	r := new(Response)
	r.route = route{StrTrim(ctrl), StrTrim(action)}
	r.redirect = true
	return r
}

// Err releases USSD session with an error message.
func (c *Context) Err(err error) *Response {
	r := new(Response)
	r.Message = err.Error()
	r.err = err
	return r
}
