package screen

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/kojogadget/kojoeditor/internal/ansi"
	"github.com/kojogadget/kojoeditor/internal/editor"
	"golang.org/x/term"
)

type Screen struct{
    width, height   int
    editor	    *editor.Editor
    originalState   *term.State
}

func NewScreen(editor *editor.Editor) *Screen {
    width, height, _ := term.GetSize(int(os.Stdout.Fd()))

    return &Screen{
	width: width, 
	height: height,
	editor: editor,
    }
}

func (s *Screen) Init() error {
    var err error
    s.originalState, err = term.MakeRaw(int(os.Stdin.Fd()))
    if err != nil {
	return err
    }

    ansi.TermHide()
    return nil
}

func (s *Screen) Close() {
    ansi.TermReset()

    if s.originalState != nil {
        term.Restore(int(os.Stdin.Fd()), s.originalState)
    }
}

func (s *Screen) GetSize() (int, int) {
    return s.width, s.height
}

func (s *Screen) Clear() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func (s *Screen) Refresh() {
    s.Clear()
    s.Render()
}

func (s *Screen) Render() {
    bufferLines := s.editor.String()
    lines := splitLines(bufferLines)

    for y, line := range lines {
	if y >= s.height {
	    break
	}
	fmt.Println(line)
    }

    cursorRow, cursorCol := s.editor.CursorPos()
    ansi.PlaseCursor(cursorRow, cursorCol)
}

func splitLines(text string) []string {
    var lines []string
    scanner := bufio.NewScanner(strings.NewReader(text))
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    
    if err := scanner.Err(); err != nil {
	fmt.Println("Error reading string:", err)
    }

    return lines
}

func (s *Screen) HandleEvents() {
    reader := bufio.NewReader(os.Stdin)
    for {
        input, _ := reader.ReadByte()
        switch input {
        case 17: // Ctrl-Q
            s.editor.HandleInput(rune(input))
            return
        case 21: // Ctrl-U
            s.editor.HandleInput(rune(input))
        case 4: // Ctrl-D
            s.editor.HandleInput(rune(input))
        case '\r', '\n':
            s.editor.HandleInput('\n')
        default:
            s.editor.HandleInput(rune(input))
        }
        s.Refresh()
    }
}
