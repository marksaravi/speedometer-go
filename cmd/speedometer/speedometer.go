package main

import (
	"fmt"

	"github.com/marksaravi/fonts-go/fonts"
)

func main() {
	var font fonts.BitmapFont = fonts.FreeMonoBold9pt7b
	fmt.Println(font.Glyphs[30].BitmapOffset)
}
