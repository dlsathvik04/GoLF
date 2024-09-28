package algorithms

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"sync"

	"github.com/dlsathvik04/GoLF/internal/server"
)

type WeightedRoundRobin struct {
	mu         sync.Mutex
	current    int
	servers    []*server.Server
	capacities []int
}

func NewWeightedRoundRobin(serverList []*server.Server, capacities []int) (*WeightedRoundRobin, error) {
	if len(serverList) != len(capacities) {
		return nil, errors.New("length of servers and capacities does not match")
	}

	if len(serverList) < 1 {
		return nil, errors.New("zero Server Exception")
	}

	field := [100]*server.Server{}

	curr := 0

	for i, cap := range capacities {
		for range cap {
			if curr > 99 {
				return nil, errors.New("invalid capacity configuration")
			}
			field[curr] = serverList[i]
			curr++
		}
	}

	if curr > 101 {
		return nil, errors.New("invalid capacity configuration")
	}

	for i := range field {
		j := rand.Intn(i + 1)
		field[i], field[j] = field[j], field[i]
	}

	return &WeightedRoundRobin{
		current:    0,
		servers:    field[:],
		capacities: capacities,
	}, nil
}

func (rr *WeightedRoundRobin) GetNextServer(_ *http.Request) (*server.Server, error) {
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
