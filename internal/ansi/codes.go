package ansi

import "fmt"

// Hide screen and store term state
func TermHide() {
    fmt.Print("\033[?1049h\033[H")
}

// Reseting the terminal
func TermReset() {
    fmt.Print("\033[?1049l")
}

// Plasing the cursor
//
// Parameters:
//      - height: The height position of the cursor (int)
//      - width: The width position of the cursor (int)
func PlaseCursor(height int, width int) {
    fmt.Printf("\033[%d;%dH", height + 1, width + 1)
}

func PrintRune(height int, width int, r rune) {
        fmt.Printf("\033[%d;%dH%c", height + 1, width + 1, r)
}

func HideCursor() {
    fmt.Print("\033[?25l")
}

func ShowCursor() {
    fmt.Print("\033[?25h")
}

