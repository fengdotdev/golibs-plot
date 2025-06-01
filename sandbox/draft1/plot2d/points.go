package plot2d

type PointDTO struct {
	ID string  `json:"id"` // Unique identifier for the point
	X  float64 `json:"x"`  // X coordinate
	Y  float64 `json:"y"`  // Y coordinate
}
type Point struct {
	id string  // Unique identifier for the point
	x  float64 // X coordinate
	y  float64 // Y coordinate
}
