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
	Lat    float64
	Lon    float64
	RadLat float64
	RadLon float64
}

type WasteFacility struct {
	Title      string
	Address    string
	WasteTypes map[string]struct{}
}
