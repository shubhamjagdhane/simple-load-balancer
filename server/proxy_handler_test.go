package server

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/shubhamjagdhane/simple-load-balancer/entity"
)

func TestIsActiveServer(t *testing.T) {
	type param struct {
		proxy *entity.ServerPool
	}
	tests := []struct {
		name string
		arg  param
		want bool
	}{
		{
			name: "Success: Returns true if server is active",
			arg: param{
				proxy: &entity.ServerPool{
					Url: "http://google.com",
				},
			},
			want: true,
		},
		{
			name: "Failure: Returns false for empty string",
			arg: param{
				proxy: &entity.ServerPool{
					Url: "",
				},
			},
			want: false,
		},
		{
			name: "Failure: Returns false for invalid URL",
			arg: param{
				proxy: &entity.ServerPool{
					Url: "abcd1029",
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				res := isActiveServer(tt.arg.proxy)
				assert.Equal(t, res, tt.want)
			},
		)
	}

}

func TestSetHealthyFlag(t *testing.T) {
	currentTime := time.Now()

	tracking := make(map[string][]time.Time)
	tracking["server-1"] = []time.Time{
		currentTime.Add(-time.Second * 10),
		currentTime.Add(-time.Second * 6),
		currentTime.Add(-time.Second * 4),
	}

	type param struct {
		proxy *entity.ServerPool
	}
	tests := []struct {
		name string
		arg  param
		want bool
	}{
		{
			name: "Failure: Will not setup Health if server is not access at least three time",
			arg: param{
				proxy: &entity.ServerPool{
					Tracking: tracking,
					Health:   false,
				},
			},
			want: false,
		},
		{
			name: "Success: Update the Health of server based on the last previous call within 15 seconds",
			arg: param{
				proxy: &entity.ServerPool{
					Tracking: tracking,
					Name:     "server-1",
					Health:   false,
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				setHealthyFlag(tt.arg.proxy)
				assert.Equal(t, tt.arg.proxy.Health, tt.want)
			},
		)
	}

}

func TestIsEmptyString(t *testing.T) {
	type param struct {
		str string
	}
	tests := []struct {
		name string
		arg  param
		want bool
	}{
		{
			name: "Succes: Should return true for empty string",
			arg: param{
				str: "",
			},
			want: true,
		},
		{
			name: "Failure: Should return false for non-empty string",
			arg: param{
				str: "some string",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				res := isEmptyString(tt.arg.str)
				assert.Equal(t, res, tt.want)
			},
		)
	}

}
