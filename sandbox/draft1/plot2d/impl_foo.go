package plot2d

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"os"
)

func (p *Plot2D) AddPoint(id string, x, y float64) error {
	p.points = append(p.points, Point{
		id: id,
		x:  x,
		y:  y,
	})
	return nil
}
func (p *Plot2D) Draw() error {

	width := p.size.Width
	height := p.size.Height
	backgroundColor := color.RGBA{255, 255, 255, 255} // Blanco opaco
	pointColor := color.RGBA{0, 0, 0, 255}
	fileName := "plot2d.png"

	img := image.NewRGBA(image.Rect(0, 0, p.size.Width, p.size.Height))


	// draw point
	centerX := width / 2
	centerY := height / 2

	err := DrawPoint(img, centerX, centerY, pointColor, 5)
	

	// Guardar la imagen como PNG
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

func DrawPoint(img *image.RGBA, x, y int, c color.Color, tickness int) error {
	if x < 0 || x >= img.Bounds().Dx() || y < 0 || y >= img.Bounds().Dy() {
		return errors.New("point coordinates out of bounds")
	}
	for dx := -tickness; dx <= tickness; dx++ {
		for dy := -tickness; dy <= tickness; dy++ {
			if x+dx < 0 || x+dx >= img.Bounds().Dx() || y+dy < 0 || y+dy >= img.Bounds().Dy() {
				continue // Skip out of bounds pixels
			}
			img.Set(x+dx, y+dy, c)
		}
	}
	return nil
}

func DrawLine(img *image.RGBA, x1, y1, x2, y2 int, c color.Color) error {
	if x1 < 0 || x1 >= img.Bounds().Dx() || y1 < 0 || y1 >= img.Bounds().Dy() ||
		x2 < 0 || x2 >= img.Bounds().Dx() || y2 < 0 || y2 >= img.Bounds().Dy() {
		return errors.New("line coordinates out of bounds")
	}

	// dx := abs(x2 - x1)
	// dy := abs(y2 - y1)
	// sx := 1
	// if x1 >= x2 {
	// 	sx = -1
	// }
	// sy := 1
	// if y1 >= y2 {
	// 	sy = -1
	// }

	err := DrawPoint(img, x1, y1, c, 5)
	if err != nil {
		return err
	}

	// if dx > dy {
	// 	err = DrawLineHorizontal(img, x1, y1, x2, c)
	// } else {
	// 	err = DrawLineVertical(img, x1, y1, y2, c)
	// }
	return err
}

func abs(i int) any {
	if i < 0 {
		return -i
	}
	return i
}
