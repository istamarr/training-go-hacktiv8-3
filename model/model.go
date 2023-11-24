package model

type SensorData struct {
	ID    int     `json:"id"`
	Water float64 `json:"water"`
	Wind  float64 `json:"wind"`
}
