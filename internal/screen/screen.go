package screen

import (
	"bufio"
	"os"
	"os/exec"

	"golang.org/x/term"
)

type Cell struct {
    Rune rune
}

type Screen struct{
    width, height   int
    cells	    [][]Cell
    cursorRow	    int
    cursorCol	    int
}

func NewScreen() *Screen {
    width, height, _ := term.GetSize(int(os.Stdout.Fd()))
    cells := make([][]Cell, height)
    for i := range cells {
	cells[i] = make([]Cell, width)
    }

    return &Screen{
	width: width, 
	height: height,
	cells: cells,
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

func (s *Screen) Render() {
    // Make the screen
}

func (s *Screen) updateCells() {
    // Update the cells on the screen
}

func (s *Screen) ReadInput() rune {
    reader := bufio.NewReader(os.Stdin)
    char, _, err := reader.ReadRune()
    if err != nil {
        panic(err)
    }
    return char
}
