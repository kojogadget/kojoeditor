package main

import (
	"os"

	"github.com/kojogadget/kojoeditor/internal/editor"
	"github.com/kojogadget/kojoeditor/internal/screen"
	"golang.org/x/term"
)

func main() {
    width, height, _ := term.GetSize(int(os.Stdout.Fd()))
    editor := editor.NewEditor(width, height)
    screen := screen.NewScreen(editor)

    if err := screen.Init(); err != nil {
	panic(err)
    }
    defer screen.Close()

    screen.HandleEvents()
}
