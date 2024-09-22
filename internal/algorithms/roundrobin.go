package algorithms

import (
	"errors"
	"fmt"
	"sync"

	"github.com/dlsathvik04/GoLF/internal/server"
)

type RoundRobin struct {
	mu      sync.Mutex
	current int
	servers []*server.Server
}

func NewRoundRobin(serverList []*server.Server) (*RoundRobin, error) {
	fmt.Println("Making a round robin algorithm implementer")
	if len(serverList) < 1 {
		return nil, errors.New("zero Server Exception")
	}
	return &RoundRobin{
		current: 0,
		servers: serverList,
	}, nil
}

func (rr *RoundRobin) GetNextServer() (*server.Server, error) {
	rr.mu.Lock()
	defer rr.mu.Unlock()
	curr := rr.current
	numServers := len(rr.servers)

	for i := 0; i < numServers; i++ {
		curr_server := (curr + i) % numServers
		fmt.Println(curr_server)
		if rr.servers[curr_server].IsHealthy {
			rr.current = (curr_server + 1) % numServers
			fmt.Println(rr.servers[curr_server])
			return rr.servers[curr_server], nil
		}
	}

	return nil, errors.New("no Healthy server error")
}
