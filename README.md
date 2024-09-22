# GoLF

GoLF (Go-Load Balancer-Firewall) is a simple implementation of a Load Balancer and a Web Application Firewall (WAF) in GoLang.

## Configuration
To configure the load balancer use the config.json file

the field ```lbAlgorithm``` stands for Load balancing algorithm it supports the following values:
```
0 - Round Robin Algorithm
1 - Weighted Round Robin Algorithm (To be implemented)
    (for this the configuration field should be an array of integers representing the percentage of requests to be sent
    the sum of all the elements should be 100.)
2 - Hashed IP algorithm (To be implemented)
```

A sample configuration for load balancing is given below.

```
{
    "servers" : [
        "http://127.0.0.1:8000",
        "http://127.0.0.1:8001"
    ],

    "capacities" : [
        67,
        33
    ],

    
    "lbAlgorithm" : 0, 

    "port" : ":9000",

    "healthCheckInterval" : "1s"
}
```

## Development Progress

- [ ] Load Balancing
    - [X] Round Robin
    - [X] Weighted Round Robin
    - [ ] Hashed IP 
- [ ] Firewall
