package banner

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"

	"github.com/urans/bing/pkg/color"
)

const (
	logo = `
██████  ██ ███    ██  ██████ 
██   ██ ██ ████   ██ ██      
██████  ██ ██ ██  ██ ██   ███ 
██   ██ ██ ██  ██ ██ ██    ██ 
██████  ██ ██   ████  ██████  

`
)

// Display show app logo
func Display() {
	s := color.Cyan.Bold(logo)
	fmt.Print(s)
}

// Print show app logo with ASCII chars
func Print() {
	figure.NewColorFigure("Bing", "", "cyan", true).Print()
}
