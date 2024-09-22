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
		currServer := (curr + i) % numServers
		fmt.Println(currServer)
		if rr.servers[currServer].IsHealthy {
			rr.current = (currServer + 1) % numServers
			fmt.Println(rr.servers[currServer])
			return rr.servers[currServer], nil
		}
	}

	return nil, errors.New("no Healthy server error")
}
