// GIF composition and encoding.
package grsys

import (
	"errors"
	"image"
	"image/gif"
	"os"
)

var encoding gif.GIF

func encode(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		Error(err)
	}
	if encoding.Image == nil {
		err = gif.Encode(f, canvas, nil)
	} else {
		Frame(500)
		err = gif.EncodeAll(f, &encoding)
	}
	if err != nil {
		f.Close()
		Error(err)
	}
	if err := f.Close(); err != nil {
		Error(err)
	}
}

func Frame(delay int) {
	if delay < 0 {
		Error(errors.New("Invalid delay"))
	}
	dup := image.NewPaletted(canvas.Rect, canvas.Palette)
	copy(dup.Pix, canvas.Pix)
	encoding.Image = append(encoding.Image, dup)
	encoding.Delay = append(encoding.Delay, delay)
}
