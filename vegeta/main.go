package main

import (
	"fmt"
	"time"

	bonehttprouter "github.com/ariefrahmansyah/bone-httprouter"
	vegeta "github.com/tsenart/vegeta/lib"
)

var rate = uint64(100) // per second
var duration = 5 * time.Second

func main() {
	var totalBone, totalHttprouter time.Duration
	for _, route := range bonehttprouter.StaticRoutes {
		totalBone += attack("http://localhost:8080", route)
		totalHttprouter += attack("http://localhost:8081", route)
	}

	fmt.Printf("Total bone: %s", totalBone)
	fmt.Printf("Total httpRouter: %s", totalHttprouter)

	fmt.Printf("Avg bone: %s", time.Duration(totalBone)/time.Duration(len(bonehttprouter.StaticRoutes)))
	fmt.Printf("Avg httpRouter: %s", time.Duration(totalHttprouter)/time.Duration(len(bonehttprouter.StaticRoutes)))
}

func attack(host string, route bonehttprouter.Route) time.Duration {
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: route.Method,
		URL:    host + route.Path,
	})

	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "bone") {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("%s: %s\n", host+route.Path, metrics.Latencies.P99)
	return metrics.Latencies.P99
}
