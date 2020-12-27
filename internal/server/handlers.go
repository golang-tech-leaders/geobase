package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"geobase/internal/model"

	"github.com/gorilla/mux"
)

func (s *Server) getLocURLForWasteType(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(s.timeout)*time.Second)
	defer func() {
		s.log.Debug().
			Str("package", "server").
			Str("func", "getLocURLForWasteType").
			Msg("canceling context")
		cancel()
	}()
	vars := mux.Vars(r)

	wasteTypeID := strings.ToLower(vars["type_id"])

	wasteTypeID, latitude, longitude, radius, err := s.getParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	recyclingPointRequest := model.RecyclingPointRequest{
		WasteTypeID: wasteTypeID,
		Longitude:   longitude,
		Latitude:    latitude,
		Radius:      int(radius),
	}

	locationForWasteType, err := s.urlFinder.GetLocationURLForWasteType(ctx, recyclingPointRequest)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			http.Error(w, "No recycling point was found for waste type: "+wasteTypeID, http.StatusNotFound)
			return
		}
	}

	result := model.MapReference{Url: locationForWasteType.Url}

	err = json.NewEncoder(w).Encode(&result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) getLocPointForWasteType(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(s.timeout)*time.Second)
	defer func() {
		s.log.Debug().
			Str("package", "server").
			Str("func", "getLocPointForWasteType").
			Msg("canceling context")
		cancel()
	}()

	wasteTypeID, latitude, longitude, radius, err := s.getParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	recyclingPointRequest := model.RecyclingPointRequest{
		WasteTypeID: wasteTypeID,
		Longitude:   longitude,
		Latitude:    latitude,
		Radius:      int(radius),
	}

	locations, err := s.locFinder.GetNearestWasteLocation(ctx, recyclingPointRequest)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			http.Error(w, "No recycling point was found for waste type: "+wasteTypeID, http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	s.log.Debug().
		Int("locations", len(locations)).
		Msg("found")

	loc := locations[0]

	result := model.LocationResponse{
		Latitude:  loc.Latitude,
		Longitude: loc.Longitude,
	}

	err = json.NewEncoder(w).Encode(&result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) getLocPointListForWasteType(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(s.timeout)*time.Second)
	defer func() {
		s.log.Debug().
			Str("package", "server").
			Str("func", "getLocPointForWasteType").
			Msg("canceling context")
		cancel()
	}()

	wasteTypeID, latitude, longitude, radius, err := s.getParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	recyclingPointRequest := model.RecyclingPointRequest{
		WasteTypeID: wasteTypeID,
		Longitude:   longitude,
		Latitude:    latitude,
		Radius:      int(radius),
	}

	locations, err := s.locFinder.GetNearestWasteLocation(ctx, recyclingPointRequest)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			http.Error(w, "No recycling point was found for waste type: "+wasteTypeID, http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	s.log.Debug().
		Int("locations", len(locations)).
		Msg("found")

	result := make([]model.LocationResponse, 0, len(locations))

	for idx := range locations {
		result = append(result, model.LocationResponse{
			Latitude:  locations[idx].Latitude,
			Longitude: locations[idx].Longitude,
		})
	}

	err = json.NewEncoder(w).Encode(&result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) getParams(r *http.Request) (wasteTypeID string, latitude, longitude float64, radius int64, err error) {
	paramErr := func(param string, cause error) error {
		return fmt.Errorf("unable to parse param %q: %s", param, err)
	}

	vars := mux.Vars(r)
	wasteTypeID = strings.ToLower(vars["type_id"])

	latitudeParam := r.URL.Query().Get("latitude")
	latitude, err = strconv.ParseFloat(latitudeParam, 64)
	if err != nil {
		err = paramErr("latitude", err)
		return
	}

	longitudeParam := r.URL.Query().Get("longitude")
	longitude, err = strconv.ParseFloat(longitudeParam, 64)
	if err != nil {
		err = paramErr("longitude", err)
		return
	}

	radiusParam := r.URL.Query().Get("radius")
	radius, err = strconv.ParseInt(radiusParam, 10, 32)
	if err != nil {
		err = paramErr("radius", err)
		return
	}

	return
}
