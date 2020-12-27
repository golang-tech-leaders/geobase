package model

type RecyclingPointRequest struct {
	WasteTypeID string
	Longitude   float32
	Latitude    float32
	Radius      int
}

type MapReference struct {
	Url string `json:"url"`
}

type RecyclingPointDBEntry struct {
	WasteType string
	Url       string
}
