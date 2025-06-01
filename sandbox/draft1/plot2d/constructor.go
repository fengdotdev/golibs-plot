package plot2d

func NewPlot2D(size Size) *Plot2D {
	return &Plot2D{
		size:   size,
		points: make([]Point, 0),
	}
}
