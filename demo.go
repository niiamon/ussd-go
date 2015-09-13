package ussd

import (
	"fmt"
	"log"
)

type demo struct {
}

func (d demo) Menu(c *Context) Response {
	menu := NewMenu()
	menu.Header = "Welcome"
	menu.AddItem("Greet me", "demo", "GreetMeForm")
	menu.AddZeroItem("Exit", "demo", "Exit")
	log.Printf("demo: %+v\n", d)
	return c.RenderMenu(menu)
}

func (d demo) GreetMeForm(c *Context) Response {
	form := NewForm("Greet Me")
	form.AddInput("Name", StrEmpty)
	form.AddInput("Sex", StrEmpty,
		NewOption("M", "Male"),
		NewOption("F", "Female"))
	return c.RenderForm(form, "demo", "GreetMe")
}

func (d demo) GreetMe(c *Context) Response {
	prefix := StrEmpty
	if c.FormData["Sex"] == "M" {
		prefix = "Master"
	} else {
		prefix = "Madam"
	}
	msg := fmt.Sprintf("%v %v"+StrNewLine, prefix, c.FormData["Name"])
	return c.Release(msg)
}

func (d *demo) Exit(c *Context) Response {
	return c.Release("Bye bye.")
}

func addData(key string, value interface{}) Middleware {
	return func(c *Context) {
		c.Data[key] = value
	}
}
