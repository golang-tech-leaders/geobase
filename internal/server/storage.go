package server

import (
	"context"
	"geobase/internal/model"
)

// URLFinder interface describes storage contract
type URLFinder interface {
	GetLocationURLForWasteType(
		ctx context.Context, request model.RecyclingPointRequest) (*model.RecyclingPointDBEntry, error)
}

// LocationFinder return a list of locations near location
type LocationFinder interface {
	GetNearestWasteLocation(ctx context.Context, request model.RecyclingPointRequest) ([]model.WasteFacility, error)
}
