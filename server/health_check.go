package server

/*
import (
	"fmt"
	"net/http"
	"time"
)

func (s *Server) HealthCheck() {
	for {
		for _, server := range s.serverPool {
			resp, err := http.Head(server.Url)
			if err != nil || resp.StatusCode != http.StatusOK {
				server.Health = false
			}

			duration := time.Since(server.LastHit)
			if server.TimesHit > 2 && duration > (15*time.Second) {
				server.Health = true
			}
			fmt.Printf("\n%v is healthy? %v\n", server.Name, server.Health)
		}
		if len(s.serverPool) == 0 {
			fmt.Printf("\nServer not found, please register the server")
		}

		time.Sleep(2 * time.Second)
	}
}*/
