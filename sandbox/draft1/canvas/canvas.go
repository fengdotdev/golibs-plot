package canvas

func Canvas1080p() *Canvas {
	return &Canvas{
		Width:  1920,
		Height: 1080,
	}
}

func Canvas720p() *Canvas {
	return &Canvas{
		Width:  1280,
		Height: 720,
	}
}

func Canvas4k() *Canvas {
	return &Canvas{
		Width:  3840,
		Height: 2160,
	}
}

func Canvas16k() *Canvas {
	return &Canvas{
		Width:  15360,
		Height: 8640,
	}
}

func NewCanvas(width, height int) *Canvas {
	return &Canvas{
		Width:  width,
		Height: height,
	}
}

type Canvas struct {
	Width  int
	Height int
	layers []Layer //[0] background, [1] foreground
}

type Layer struct {
	Name    string
	Width   int
	Height  int
	Visible bool
}
