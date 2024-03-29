package models

type PlayKnightCardRequest struct {
	DevelopmentCardID string `json:"developmentCardID" validate:"required,objectid"`
	TerrainID         string `json:"terrainID" validate:"required,objectid"`
	PlayerID          string `json:"playerID" validate:"omitempty,objectid"`
}
