package prm

import (
	"fmt"
	"os/exec"
)

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
	cmd := exec.Command("primitive", "-i", inPath, "-o", outPath, "-m", fmt.Sprint(mode), "-n", fmt.Sprint(shapes))
	return cmd.Run()
}
