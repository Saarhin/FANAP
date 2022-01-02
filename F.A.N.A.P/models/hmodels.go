package models

type Inputmodel struct {
	Main  Rectangle   `json:"main"`
	Input []Rectangle `json:"input"`
}

type Point struct {
	XR float64 `json:"xr"`
	YR float64 `json:"yr"`
}
