package plot2d

import "image"

func (p *Plot2D) AddPoint(id string, x, y float64) error {
	p.points = append(p.points, Point{
		id: id,
		x:  x,
		y:  y,
	})
	return nil
}
func (p *Plot2D) Draw() {
	img := image.NewRGBA(image.Rect(0, 0, p.size.Width, p.size.Height))

	// 
}
