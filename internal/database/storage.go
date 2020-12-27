package database

import (
	"context"
	"encoding/json"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"geobase/internal/model"
)

// Storage of app
type Storage struct {
	Locations      map[model.Coordinate]model.WasteFacility
	metersInRadius int
}

type WastePlace struct {
	ID        string `json:"id"`
	Latitude  string `json:"lat"`
	Longitude string `json:"lng"`
	Title     string `json:"title"`
	Address   string `json:"address"`
	Content   string `json:"content_text"`
}

// NewLocationFinder creates a new server
func NewLocationFinder(datapath string, metersInRadius int) (*Storage, error) {
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
		Locations:      make(map[model.Coordinate]model.WasteFacility, len(locations)),
		metersInRadius: metersInRadius,
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

		s.Locations[model.Coordinate{
			Lat:    lat,
			Lon:    lon,
			RadLat: degreesToRadians(lat),
			RadLon: degreesToRadians(lon),
		}] = wf
	}

	return &s, nil
}

// GetNearestWasteLocation returns a list of nearest locations
func (s Storage) GetNearestWasteLocation(
	ctx context.Context, req model.RecyclingPointRequest) ([]model.Coordinate, error) {

	locations := s.linearSearch(ctx, req.Latitude, req.Longitude, float64(req.Radius), req.WasteTypeID)

	if len(locations) == 0 {
		return nil, model.ErrNotFound
	}

	sort.SliceStable(locations, func(i, j int) bool {
		return locations[i].radius < locations[j].radius
	})

	res := make([]model.Coordinate, 0, len(locations))
	for i := range locations {
		res = append(res, locations[i].coords)
	}

	return res, nil
}

// linearSearch does suboptimal linear search
func (s Storage) linearSearch(ctx context.Context,
	lat, lon, rad float64, wType string) (res []foundLocation) {
	var (
		radLat = degreesToRadians(lat)
		radLon = degreesToRadians(lon)
		radius = rad * float64(s.metersInRadius)
	)

	for c, l := range s.Locations {
		select {
		case <-ctx.Done():
			return
		default:
		}

		dist := distance(radLat, radLon, c.RadLat, c.RadLon)

		if dist < radius {
			if _, ok := l.WasteTypes[wType]; ok {
				res = append(res, foundLocation{
					radius: dist,
					coords: c,
				})
			}
		}
	}
	return
}

// degreesToRadians converts from degrees to radians.
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

// distance finds distance between two points in meters
func distance(lat1, lon1, lat2, lon2 float64) float64 {
	const r = 6378100 // Earth radius in METERS
	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	return 2 * r * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}

// transformWasteType transforms free input type
func transformWasteType(t string) string {
	t = strings.Replace(t, " ", "", -1)
	t = strings.ToLower(t)
	return t
}

type foundLocation struct {
	radius float64
	coords model.Coordinate
}
