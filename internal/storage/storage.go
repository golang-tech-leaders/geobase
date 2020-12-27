package storage

import (
	"context"
	"encoding/json"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"geobase/internal/model"
)

// Storage of app
type Storage struct {
	Locations map[model.Coordinate]model.WasteFacility
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
		Locations: make(map[model.Coordinate]model.WasteFacility, len(locations)),
	}

	for _, l := range locations {
		types := strings.Split(l.Content, ",")
		wf := model.WasteFacility{
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

		wf.Coordinate = model.Coordinate{
			Latitude:  lat,
			Longitude: lon,
		}

		s.Locations[wf.Coordinate] = wf
	}

	return &s, nil
}

// GetNearestWasteLocation returns a list of nearest locations
func (s Storage) GetNearestWasteLocation(
	ctx context.Context, req model.RecyclingPointRequest) ([]model.WasteFacility, error) {

	locations := s.linearSearch(ctx, req.Latitude, req.Longitude, float64(req.Radius))

	validLocations := make([]model.WasteFacility, 0)
	for idx := range locations {
		if _, ok := locations[idx].WasteTypes[req.WasteTypeID]; ok {
			validLocations = append(validLocations, locations[idx])
		}
	}

	if len(validLocations) == 0 {
		return nil, model.ErrNotFound
	}

	return validLocations, nil
}

// linearSearch does suboptimal linear search
func (s Storage) linearSearch(ctx context.Context,
	lon, lat, rad float64) (res []model.WasteFacility) {
	for c, l := range s.Locations {
		select {
		case <-ctx.Done():
			return
		default:
		}

		if math.Hypot(c.Longitude-lon, c.Latitude-lat) < rad {
			res = append(res, l)
		}
	}
	return
}

// transformWasteType transforms free input type
func transformWasteType(t string) string {
	t = strings.Replace(t, " ", "", -1)
	t = strings.ToLower(t)
	return t
}
