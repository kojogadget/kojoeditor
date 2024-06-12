package main

import (
	"fmt"
	"os"

	"github.com/kojogadget/kojoeditor/internal/editor"
	"github.com/kojogadget/kojoeditor/internal/term"
	"github.com/kojogadget/kojoeditor/internal/screen"
)

func main() {
    if err := term.EnableRawMode(); err != nil {
        fmt.Fprintln(os.Stderr, "Unable to enter raw mode:", err)
        os.Exit(1)
    }
    defer term.DisableRawMode()

    screen := screen.NewScreen()
    e := editor.NewEditor(screen.GetSize())

    screen.Render()

    for {
	input := screen.ReadInput()
	e.HandleInput(input)
	screen.Render()
	if e.ShouldExit() {
	    break
	}
    }
}
