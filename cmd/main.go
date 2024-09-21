package main

import (
	"fmt"
	"loadbalancer/internal/config"
	"loadbalancer/internal/loadbalancer"
	"net/http"
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
