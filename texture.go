package sprite

import (
	"github.com/go-gl/gl"
	"image"
	_ "image/png"
	"log"
	"os"
)

func Upload(i *image.NRGBA, tex gl.Texture, x, y, z int) Image {
	width := i.Rect.Max.X
	height := i.Rect.Max.Y

	tex.Bind(gl.TEXTURE_2D_ARRAY)
	gl.TexSubImage3D(gl.TEXTURE_2D_ARRAY, 0, x, y, z, width, height, 1, gl.RGBA, gl.UNSIGNED_BYTE, i.Pix)
	return Image{
		TextureLeft:   float32(x),
		TextureTop:    float32(y),
		TextureRight:  float32(x + width),
		TextureBottom: float32(y + height),
		Layer:         float32(z)}
}

func LoadPNG(filename string) image.Image {
	// avataan tiedosto
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// tehdään tiedostosta kuva
	kuva, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return kuva
}
