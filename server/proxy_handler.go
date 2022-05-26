package server

import (
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/shubhamjagdhane/simple-load-balancer/entity"
	"github.com/shubhamjagdhane/simple-load-balancer/errors"
)

func (s *Server) ProxyHandler(w http.ResponseWriter, req *http.Request) error {
	if selectedServer := s.getHealthyServer(); selectedServer != nil {
		selectedServer.ServeHTTP(w, req)
		w.WriteHeader(http.StatusOK)
		return nil
	}
	return errors.GetError("No Healthy Server Available", http.StatusServiceUnavailable)
}

func (s *Server) getHealthyServer() *httputil.ReverseProxy {
	if totalServers := uint(len(s.serverPool)); totalServers > 0 {
		counter := uint(0)
		currentIndex := s.selectedServerIndex

		for counter != totalServers {
			counter += 1
			proxy := s.serverPool[currentIndex]

			setHealthyFlag(proxy)
			if isActiveServer(proxy) {
				updateTracking(proxy)
				s.serverPool[currentIndex] = proxy
				s.selectedServerIndex = (currentIndex + 1) % totalServers
				return proxy.ReverseProxy
			}

			currentIndex = (currentIndex + 1) % totalServers
		}
		return nil
	}
	return nil
}

func updateTracking(proxy *entity.ServerPool) {
	if successTimeSlice, ok := proxy.Tracking[proxy.Name]; ok {
		if len(successTimeSlice) == 3 {
			successTimeSlice = successTimeSlice[1:]
		}
		successTimeSlice = append(successTimeSlice, time.Now())
		proxy.Tracking[proxy.Name] = successTimeSlice
	} else {
		proxy.Tracking[proxy.Name] = []time.Time{time.Now()}
	}

}

func isActiveServer(proxy *entity.ServerPool) bool {
	res, err := http.Head(proxy.Url)
	if err != nil || res.StatusCode != http.StatusOK {
		return false
	}
	return true
}

func setHealthyFlag(proxy *entity.ServerPool) {
	if track := proxy.Tracking[proxy.Name]; len(track) == 3 {
		first := track[0]
		if time.Since(first) <= (time.Second * 15) {
			proxy.Health = true
		} else {
			proxy.Health = false
		}
	}
}
