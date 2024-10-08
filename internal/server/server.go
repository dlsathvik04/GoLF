package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

type Server struct {
	URL                 *url.URL
	IsHealthy           bool
	proxy               *httputil.ReverseProxy
	healthCheckInterval time.Duration
	mu                  sync.Mutex
}

func NewServer(address string, healthCheckInterval time.Duration) (*Server, error) {
	serverURL, err := url.Parse(address)
	if err != nil {
		return nil, err
	}
	reverseProxy := httputil.NewSingleHostReverseProxy(serverURL)
	server := &Server{
		URL:                 serverURL,
		IsHealthy:           true,
		proxy:               reverseProxy,
		healthCheckInterval: healthCheckInterval,
	}

	go server.CheckHealth()
	return server, nil
}

func (s *Server) CheckHealth() {
	fmt.Println("Starting health loop for : ", s.URL.String(), s.healthCheckInterval)
	for range time.Tick(s.healthCheckInterval) {
		res, err := http.Head(s.URL.String())
		s.mu.Lock()
		if err != nil || res.StatusCode != http.StatusOK {
			s.IsHealthy = false
		} else {
			s.IsHealthy = true
		}
		s.mu.Unlock()
		fmt.Println("Server : ", s.URL.String(), " Health : ", s.IsHealthy)
	}

}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a request to server: ", s.URL.String(), r.RequestURI, r.Method)
	w.Header().Add("X-Forwarded-Server", s.URL.String())
	s.proxy.ServeHTTP(w, r)
}
