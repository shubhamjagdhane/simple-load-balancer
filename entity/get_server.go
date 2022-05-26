package entity

import "time"

type GetServer struct {
	Name     string                 `json:"server_name"`
	Url      string                 `json:"url"`
	Healthy  bool                   `json:"healthy"`
	Tracking map[string][]time.Time `json:"track"`
}
