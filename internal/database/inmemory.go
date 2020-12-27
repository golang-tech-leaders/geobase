package database

import (
	"context"

	"geobase/internal/model"
)

type InmemoryDB map[string]string

func New() InmemoryDB {
	db := make(InmemoryDB, 10)
	return db
}

func (db InmemoryDB) Init() {
	db["стекло"] = "https://www.google.com/maps/@55.7324321,37.6632525,20.99z"
	db["пластик"] = "https://www.google.com/maps/@55.7324321,37.6632525,20.99z"
	db["метал"] = "https://www.google.com/maps/@55.7188514,37.574609,17.97z"
	db["одежда"] = "https://www.google.com/maps/@55.7188514,37.574609,17.97z"
	db["иное"] = "https://www.google.com/maps/@55.7188514,37.574609,17.97z"
	db["опасное"] = "https://www.google.com/maps/@55.7324321,37.6632525,20.99z"
	db["батарейки"] = "https://www.google.com/maps/@55.7188514,37.574609,17.97z"
	db["лампочки"] = "https://www.google.com/maps/@55.7188514,37.574609,17.97z"
	db["техника"] = "https://www.google.com/maps/@55.7188514,37.574609,17.97z"
	db["тетрапак"] = "https://www.google.com/maps/@55.7188514,37.574609,17.97z"
	db["крышечки"] = "https://www.google.com/maps/@55.7188514,37.574609,17.97z"
	db["шины"] = "https://www.google.com/maps/@55.7188514,37.574609,17.97z"
}

func (db InmemoryDB) GetLocationURLForWasteType(_ context.Context, recyclingPointRequest model.RecyclingPointRequest) (*model.RecyclingPointDBEntry, error) {
	if url, ok := db[recyclingPointRequest.WasteTypeID]; ok {
		return &model.RecyclingPointDBEntry{WasteType: recyclingPointRequest.WasteTypeID, Url: url}, nil
	}
	return nil, model.ErrNotFound
}
