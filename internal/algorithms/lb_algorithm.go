package algorithms

import (
	"log"

	"github.com/dlsathvik04/GoLF/internal/server"
)

type LBAlgorithm interface {
	GetNextServer() (*server.Server, error)
}

func NewLBAlgorithm(serverList []*server.Server, priorityList []int) LBAlgorithm {
	rr, err := NewRoundRobin(serverList)
	if err != nil {
		log.Fatal("Cannot create new round robin load balancer")
	}

	return rr
}
