package editor

import (
    "github.com/kojogadget/kojoeditor/internal/editor/rope"
)

type Editor struct {
    buffer          *rope.Rope
    shouldExit      bool
    viewWidth       int 
    viewHeight      int
    topLine         int
    cursorRow       int
    cursorCol       int
}

func NewEditor(width, height int) *Editor {
    return &Editor{
        buffer:     rope.NewRope(""),
        shouldExit: false,
        viewWidth:  width,
        viewHeight: height,
        topLine:    0,
        cursorRow:  0,
        cursorCol:  0,
    }
}

func (e *Editor) InsertRune(r rune) {
    e.buffer = e.buffer.Insert(e.cursorCol, string(r))
    e.cursorCol++
}

func (e *Editor) NewLine() {
    // Make a new line
    e.buffer = e.buffer.Insert(e.cursorCol, "\n")
    e.cursorRow++
    e.cursorCol = 0
}

func (e *Editor) ScrollUp() {
    if e.topLine > 0 {
        e.topLine--
    }
    if e.cursorRow > 0 {
        e.cursorRow--
    }
}

func (e *Editor) ScrollDown() {
    if e.topLine+e.viewHeight < len(e.buffer.String()) {
        e.topLine++
    }
    if e.cursorRow < e.viewHeight - 1 {
        e.cursorRow++
    }
}

func (e *Editor) HandleInput(input rune) {
    if input == 17 { // Ctrl-Q
        e.shouldExit = true
    } else if input == 21 { // Ctrl-U
        e.ScrollUp()
    } else if input == 4 { // Ctrl-D
        e.ScrollDown()
    } else if input == '\r' || input == '\n' {
        e.NewLine()
    } else {
        e.InsertRune(input)
    }
}

func (e *Editor) CursorPos() (int, int) {
    return e.cursorRow, e.cursorCol
}

func (e *Editor) ShouldExit() bool {
    return e.shouldExit
}

func (e *Editor) String() string {
    return e.buffer.String()
}
