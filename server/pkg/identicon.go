package pkg

import (
	"image/png"
	"os"
	"path/filepath"

	"github.com/ki365/identigo"
)

// GenerateIdenticon generates a PNG identicon with filepath fp, hash string str, and pixel size of s.
func GenerateIdenticonPNG(fp, str string, s int) error {
	img := identigo.GenerateFromString(str, s, s)
	f := fp

	if filepath.Ext(fp) != ".png" {
		f += ".png"
	}

	file, err := os.Create(f)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}
