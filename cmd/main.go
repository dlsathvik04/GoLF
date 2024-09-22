package main

import (
	"fmt"
	"net/http"

	"github.com/dlsathvik04/GoLF/internal/config"
	"github.com/dlsathvik04/GoLF/internal/loadbalancer"
)

func main() {
	conf := config.GetConfig()
	fmt.Println("Completed Reading Config: ", conf)
	fmt.Println("Building Load Balancer....")
	lb := loadbalancer.NewLoadBalancer(conf)
	fmt.Println("Built the load balancer")

	http.HandleFunc("/", lb.ServeHTTP)
	fmt.Println("listening on : ", conf.Port)
	http.ListenAndServe(conf.Port, nil)
}
