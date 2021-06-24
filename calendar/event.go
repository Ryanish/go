package calendar

import (
	"errors"
	"unicode/utf8"
)

//by changing the type param "title" to lower case, it becoms unexportable, therefore we setup getter or setter functions to encapsulate and allow validation at this level
type Event struct {
	title string
	Date
}

// Getter method
func (e *Event) Title() string {
	return e.title
}

//Setter method
func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("string too long")
	}
	e.title = title
	return nil
}
