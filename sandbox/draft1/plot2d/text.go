package plot2d

import (
	"image"
	"image/color"
	"image/draw"
	"os"


	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

// Dibuja texto en una imagen en la posición (x, y) con el color especificado
func dibujarTexto(img *image.RGBA, texto string, x, y int, textoColor color.Color, rutaFuente string, tamanoFuente float64) error {
	// Cargar la fuente TTF
	fontBytes, err := os.ReadFile(rutaFuente)
	if err != nil {
		return err
	}

	f, err := opentype.Parse(fontBytes)
	if err != nil {
		return err
	}

	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    tamanoFuente, // Tamaño de la fuente en puntos
		DPI:     72,           // Resolución DPI
		Hinting: font.HintingFull,
	})
	if err != nil {
		return err
	}
	defer face.Close()

	// Crear una capa temporal para el texto
	capaTexto := image.NewRGBA(img.Bounds())
	draw.Draw(capaTexto, capaTexto.Bounds(), &image.Uniform{color.Transparent}, image.Point{0, 0}, draw.Src)

	// Dibujar el texto en la capa temporal
	d := &font.Drawer{
		Dst:  capaTexto,
		Src:  image.NewUniform(textoColor),
		Face: face,
		Dot:  fixed.P(x, y), // Posición del texto
	}
	d.DrawString(texto)

	// Combinar la capa de texto con la imagen original
	draw.Draw(img, img.Bounds(), capaTexto, image.Point{0, 0}, draw.Over)

	return nil
}