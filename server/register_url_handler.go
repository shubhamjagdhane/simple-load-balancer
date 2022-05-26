package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/shubhamjagdhane/simple-load-balancer/entity"
	"github.com/shubhamjagdhane/simple-load-balancer/errors"
)

func (s *Server) RegisterUrlHandler(w http.ResponseWriter, req *http.Request) error {
	w.Header().Set("Content-Type", "application/json")

	if len(s.serverPool) == math.MaxInt16 {
		return errors.GetError("Too many servers, limit exceed", http.StatusBadRequest)
	}

	registerUrlReq := &entity.RegisterRequest{}
	body, err := ioutil.ReadAll(req.Body)

	defer req.Body.Close()
	if err != nil {
		return errors.GetError(errors.BadRequest, http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &registerUrlReq)
	if err != nil {
		return errors.GetError("Error while Unmarshal the registerUrlReq", http.StatusInternalServerError)
	}

	if s.isHostExist(registerUrlReq.Url) {
		return errors.GetError("Server has been already registered", http.StatusBadRequest)
	}

	u, err := url.Parse(registerUrlReq.Url)
	if err != nil {
		return errors.GetError("Error while parsing the URL", http.StatusBadRequest)
	}

	rp := httputil.NewSingleHostReverseProxy(u)
	s.serverPool = append(s.serverPool, &entity.ServerPool{
		Name:         registerUrlReq.Name,
		Url:          registerUrlReq.Url,
		Health:       registerUrlReq.Health,
		ReverseProxy: rp,
		Tracking:     make(map[string][]time.Time),
	})

	registerResponse := entity.RegisterResponse{
		Message:    fmt.Sprintf(`%v server has been successfully registered`, registerUrlReq.Name),
		StatusCode: http.StatusCreated,
	}

	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(registerResponse)
}

func (s *Server) isHostExist(hostname string) bool {
	for _, hosts := range s.serverPool {
		if strings.EqualFold(hosts.Url, hostname) {
			return true
		}
	}
	return false
}
