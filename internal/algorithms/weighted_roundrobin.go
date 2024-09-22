package algorithms

import (
	"errors"
	"sync"

	"github.com/dlsathvik04/GoLF/internal/server"
)

type WeightedRoundRobin struct {
	mu         sync.Mutex
	current    int
	servers    []*server.Server
	capacities []float32
}

func NewWeightedRoundRobin(serverList []*server.Server, capacities []float32) (*WeightedRoundRobin, error) {
	if len(serverList) != len(capacities) {
		return nil, errors.New("length of servers and capacities does not match")
	}

	if len(serverList) < 1 {
		return nil, errors.New("zero Server Exception")
	}

	return &WeightedRoundRobin{
		current:    0,
		servers:    serverList,
		capacities: capacities,
	}, nil
}

// func (wrr *WeightedRoundRobin) GetNextServer() (*server.Server, error) {
// 	randomNum := rand.Float32()

// 	var cum float32 = 0

// 	for index, weight := range wrr.capacities {
// 		cum += weight
// 		if randomNum < cum {
// 			return
// 		}
// 	}
// }
