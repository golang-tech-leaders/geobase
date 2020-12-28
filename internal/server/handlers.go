package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"geobase/internal/model"

	"github.com/gorilla/mux"
)

func (s *Server) getLocURLForWasteType(w http.ResponseWriter, r *http.Request) {
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

	locationForWasteType, err := s.urlFinder.GetLocationURLForWasteType(r.Context(), recyclingPointRequest)
	if err != nil {
		s.log.Error().
			Err(err).
			Str("wasteType", wasteTypeID).
			Float64("latitude", latitude).
			Float64("longitude", longitude).
			Str("func", "getLocURLForWasteType").
			Msg("processing failed")

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

	locations, err := s.locFinder.GetNearestWasteLocation(r.Context(), recyclingPointRequest)
	if err != nil {
		s.log.Error().
			Err(err).
			Str("wasteType", wasteTypeID).
			Float64("latitude", latitude).
			Float64("longitude", longitude).
			Str("func", "getLocPointForWasteType").
			Msg("processing failed")

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
		Latitude:  loc.Lat,
		Longitude: loc.Lon,
	}

	err = json.NewEncoder(w).Encode(&result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) getLocPointListForWasteType(w http.ResponseWriter, r *http.Request) {
	wasteTypeID, latitude, longitude, radius, err := s.getParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	recyclingPointRequest := model.RecyclingPointRequest{
		WasteTypeID: wasteTypeID,
		Latitude:    latitude,
		Longitude:   longitude,
		Radius:      int(radius),
	}

	locations, err := s.locFinder.GetNearestWasteLocation(r.Context(), recyclingPointRequest)
	if err != nil {
		s.log.Error().
			Err(err).
			Str("wasteType", wasteTypeID).
			Float64("latitude", latitude).
			Float64("longitude", longitude).
			Str("func", "getLocPointListForWasteType").
			Msg("processing failed")

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
			Latitude:  locations[idx].Lat,
			Longitude: locations[idx].Lon,
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
