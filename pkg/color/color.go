package color

import (
	"fmt"
)

// Color represents a text color.
type Color uint8

// Terminal Colors
const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// Add adds the color to the given string
func (c Color) Add(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(c), s)
}

// Any adds the color to any type
func (c Color) Any(s interface{}) string {
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", uint8(c), s)
}

// Bold adds a bold color to the given string
func (c Color) Bold(s string) string {
	return fmt.Sprintf("\x1b[1;%dm%s\x1b[0m", uint8(c), s)
}

// Print .
func (c Color) Print(s ...any) {
	fmt.Println(c.Any(s))
}

// Printf .
func (c Color) Printf(format string, s ...any) {
	fs := fmt.Sprintf(format, s)
	fmt.Println(c.Any(fs))
}
