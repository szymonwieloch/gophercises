package prm

import (
	"fmt"
	"os/exec"
)

//go:generate enumer -type=Mode
type Mode uint

const (
	Combo          Mode = 0
	Triangle       Mode = 1
	Rectangle      Mode = 2
	Ellipse        Mode = 3
	Circle         Mode = 4
	RotateDirect   Mode = 5
	Beziers        Mode = 6
	RotatedEllipse Mode = 7
	Polygon        Mode = 8
)

func Transform(inPath string, outPath string, shapes uint, mode Mode) error {
	args := []string{"-i", inPath, "-o", outPath, "-m", fmt.Sprint(uint(mode)), "-n", fmt.Sprint(shapes)}
	cmd := exec.Command("primitive", args...)
	return cmd.Run()
}
