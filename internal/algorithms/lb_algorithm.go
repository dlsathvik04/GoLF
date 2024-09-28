package algorithms

import (
	"log"
	"net/http"

	"github.com/dlsathvik04/GoLF/internal/server"
)

type LBAlgorithm interface {
	GetNextServer(req *http.Request) (*server.Server, error)
}

func NewLBAlgorithm(serverList []*server.Server, priorityList []int) LBAlgorithm {
	rr, err := NewRoundRobin(serverList)
	if err != nil {
		log.Fatal("Cannot create new round robin load balancer")
	}

	return rr
}
