package term

import (
	"os"

	"github.com/kojogadget/kojoeditor/internal/ansi"
	"golang.org/x/term"
)

var originalState *term.State

func EnableRawMode() error {
    var err error
    originalState, err = term.MakeRaw((int(os.Stdin.Fd())))
    if err != nil {
        return err
    }

    ansi.TermHide()

    return nil
}

func DisableRawMode() error {
    ansi.TermReset()

    if originalState != nil {
        return term.Restore(int(os.Stdin.Fd()), originalState)
    }
    return nil
}
