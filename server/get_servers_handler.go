package server

import (
	"encoding/json"
	"net/http"

	"github.com/shubhamjagdhane/simple-load-balancer/entity"
)

func (s *Server) GetServers(w http.ResponseWriter, req *http.Request) error {
	registeredServers := []*entity.GetServer{}
	for _, rs := range s.serverPool {
		registeredServer := &entity.GetServer{
			Name:     rs.Name,
			Url:      rs.Url,
			Healthy:  rs.Health,
			Tracking: rs.Tracking,
		}
		registeredServers = append(registeredServers, registeredServer)
	}
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(registeredServers)
}
