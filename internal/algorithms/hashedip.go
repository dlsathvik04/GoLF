package algorithms

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"unicode"

	"github.com/dlsathvik04/GoLF/internal/server"
)

type HashedIP struct {
	mu      sync.Mutex
	servers []*server.Server
}

func NewHashedIP(serverList []*server.Server) (*HashedIP, error) {
	fmt.Println("Making a round robin algorithm implementer")
	if len(serverList) < 1 {
		return nil, errors.New("zero Server Exception")
	}
	return &HashedIP{
		servers: serverList,
	}, nil
}

func getIPDigitSum(ip string) int {
	res := 0
	for _, char := range ip {
		if unicode.IsDigit(char) {
			val, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}
			res += val
		}
	}
	return res
}

func (hi *HashedIP) GetNextServer(req *http.Request) (*server.Server, error) {
	hi.mu.Lock()
	defer hi.mu.Unlock()
	curr := getIPDigitSum(req.RemoteAddr) % len(hi.servers)
	numServers := len(hi.servers)
	for i := 0; i < numServers; i++ {
		currServer := (curr + i) % numServers
		fmt.Println(currServer)
		if hi.servers[currServer].IsHealthy {
			fmt.Println(hi.servers[currServer])
			return hi.servers[currServer], nil
		}
	}

	return nil, errors.New("no Healthy server error")
}
