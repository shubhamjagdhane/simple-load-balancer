package entity

import (
	"net/http/httputil"
	"time"
)

type ServerPool struct {
	Name         string
	Url          string
	Health       bool
	ReverseProxy *httputil.ReverseProxy
	Tracking     map[string][]time.Time
}
