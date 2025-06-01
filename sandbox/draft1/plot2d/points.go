package plot2d

import "image/color"

type PointDTO struct {
	ID    string  `json:"id"`    // Unique identifier for the point
	X     float64 `json:"x"`     // X coordinate
	Y     float64 `json:"y"`     // Y coordinate
	Color string  `json:"color"` // Color in hex format (e.g., "#FF0000" for red)
}
type Point struct {
	id string     // Unique identifier for the point
	x  float64    // X coordinate
	y  float64    // Y coordinate
	pc color.RGBA // Point color
}
