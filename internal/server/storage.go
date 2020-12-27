package server

import (
	"context"
	"geobase/internal/model"
)

// GeobaseRepository interface describes storage contract
type GeobaseRepository interface {
	GetLocationForWasteType(ctx context.Context, recyclingPointRequest model.RecyclingPointRequest) (*model.RecyclingPointDBEntry, error)
}
