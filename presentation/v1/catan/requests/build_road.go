package requests

type BuildRoad struct {
	PathID string `json:"pathID" validate:"required,objectid"`
}
