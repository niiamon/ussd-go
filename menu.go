package ussd

import (
	"fmt"
)

// Menu for USSD
type Menu struct {
	Header, Footer string
	Items          []*MenuItem
	ZeroItem       *MenuItem
}

// MenuItem is an item to be displayed on USSD menu.
// Usually not used directly. Use `(*Menu).AddItem` or `(*Menu).AddZeroItem`.
type MenuItem struct {
	Name  string
	Route route
}

// NewMenu creates a new Menu
func NewMenu() *Menu {
	return &Menu{
		Items: make([]*MenuItem, 0),
	}
}

// AddItem to USSD menu.
func (m *Menu) AddItem(name, ctrl, action string) *Menu {
	item := &MenuItem{name, route{ctrl, action}}
	m.Items = append(m.Items, item)
	return m
}

// AddZeroItem adds an item at the bottom of USSD menu.
// This item always routes to a choice of "0".
func (m *Menu) AddZeroItem(name, ctrl, action string) *Menu {
	m.ZeroItem = &MenuItem{name, route{ctrl, action}}
	return m
}

// Render USSD menu.
func (m Menu) Render() string {
	msg := StrEmpty
	if m.Header != StrEmpty {
		msg += m.Header + StrNewLine
	}
	for i, item := range m.Items {
		msg += fmt.Sprintf("%d. %v"+StrNewLine, i+1, item.Name)
	}
	if m.ZeroItem != nil {
		msg += "0. " + m.ZeroItem.Name + StrNewLine
	}
	msg += m.Footer
	return msg
}
