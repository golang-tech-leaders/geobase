package model

type RecyclingPointRequest struct {
	WasteTypeID string
	Longitude   float64
	Latitude    float64
	Radius      int
}

type MapReference struct {
	Url string `json:"url"`
}

type LocationResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type RecyclingPointDBEntry struct {
	WasteType string
	Url       string
}

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

type WasteFacility struct {
	Coordinate
	Title      string
	Address    string
	WasteTypes map[string]struct{}
}
