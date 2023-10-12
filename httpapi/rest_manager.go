package httpapi

import (
	"net/http"
	"time"
)

func (s *Server) syncStats(_ *http.Request) (int, interface{}) {
	type response struct {
		LastRun time.Time `json:"last_run"`
		NextRun time.Time `json:"next_run"`
	}

	l, n := s.manager.Stats()

	return http.StatusOK, &response{LastRun: l, NextRun: n}
}

func (s *Server) syncTrigger(_ *http.Request) (int, interface{}) {
	s.manager.Trigger()

	return http.StatusOK, nil
}
