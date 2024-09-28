package loadbalancer

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dlsathvik04/GoLF/internal/algorithms"
	"github.com/dlsathvik04/GoLF/internal/config"
	"github.com/dlsathvik04/GoLF/internal/server"
)

type LoadBalancer struct {
	algorithm algorithms.LBAlgorithm
}

func NewLoadBalancer(config *config.Config) *LoadBalancer {
	var serverList []*server.Server
	for _, severAddress := range config.Servers {
		duration, err := time.ParseDuration(config.HealthCheckInterval)
		if err != nil {
			log.Fatal("Invalid Health check interval configuration")
		}
		fmt.Println(duration)

		currServer, err := server.NewServer(severAddress, duration)
		if err != nil {
			log.Fatal(err.Error())
		}
		serverList = append(serverList, currServer)
	}
	switch config.Algorithm {
	case 0:
		alg, err := algorithms.NewRoundRobin(serverList)
		if err != nil {
			log.Fatal(err.Error())
		}
		return &LoadBalancer{
			algorithm: alg,
		}
	case 1:
		alg, err := algorithms.NewWeightedRoundRobin(serverList, config.Capacities)
		if err != nil {
			log.Fatal(err.Error())
		}
		return &LoadBalancer{
			algorithm: alg,
		}
	case 2:
		alg, err := algorithms.NewHashedIP(serverList)
		if err != nil {
			log.Fatal(err.Error())
		}
		return &LoadBalancer{
			algorithm: alg,
		}
	default:
		log.Fatal("cant determine the Load Balancing algorithm")
		return nil
	}
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a  request")
	server, err := lb.algorithm.GetNextServer(r)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("error getting a new server")
	}
	server.ServeHTTP(w, r)
}
