package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"geobase/internal/models"
)

// Storage of app
type Storage struct {
	Locations map[models.Coordinate]models.WasteFacility
}

type WastePlace struct {
	ID        string `json:"id"`
	Latitude  string `json:"lat"`
	Longitude string `json:"lng"`
	Title     string `json:"title"`
	Address   string `json:"address"`
	Content   string `json:"content_text"`
}

// New creates a new server
func New(datapath string) (*Storage, error) {
	dir, _ := os.Getwd()
	f, err := os.Open(filepath.Join(dir, datapath))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var locations map[string]WastePlace
	err = json.NewDecoder(f).Decode(&locations)
	if err != nil {
		return nil, err
	}

	s := Storage{
		Locations: make(map[models.Coordinate]models.WasteFacility, len(locations)),
	}

	for _, l := range locations {
		types := strings.Split(l.Content, ",")
		wf := models.WasteFacility{
			Title:      l.Title,
			Address:    l.Address,
			WasteTypes: make(map[string]struct{}, len(types)),
		}
		for _, t := range types {
			wf.WasteTypes[transformWasteType(t)] = struct{}{}
		}

		lat, err := strconv.ParseFloat(l.Latitude, 64)
		if err != nil {
			return nil, err
		}
		lon, err := strconv.ParseFloat(l.Longitude, 64)
		if err != nil {
			return nil, err
		}

		s.Locations[models.Coordinate{
			Latitude:  lat,
			Longitude: lon,
		}] = wf
	}

	return &s, nil
}

// transformWasteType transforms free input type
func transformWasteType(t string) string {
	t = strings.TrimSpace(t)
	t = strings.ToLower(t)
	return t
}
