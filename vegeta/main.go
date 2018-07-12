package main

import (
	"fmt"
	"time"

	bonehttprouter "github.com/ariefrahmansyah/bone-httprouter"
	vegeta "github.com/tsenart/vegeta/lib"
)

var rate = uint64(1000) // per second
var duration = 10 * time.Second

func main() {
	var totalBase, totalBone, totalHttprouter, totalRouter time.Duration
	for _, route := range bonehttprouter.StaticRoutes {
		totalBase += attack("http://localhost:8080", route)
		totalBone += attack("http://localhost:8081", route)
		totalHttprouter += attack("http://localhost:8082", route)
		totalRouter += attack("http://localhost:8083", route)
		fmt.Println()
	}

	fmt.Printf("Total base: %s\n", totalBase)
	fmt.Printf("Total bone: %s\n", totalBone)
	fmt.Printf("Total httprouter: %s\n", totalHttprouter)
	fmt.Printf("Total router: %s\n", totalRouter)

	fmt.Printf("Avg base: %s\n", time.Duration(totalBase)/time.Duration(len(bonehttprouter.StaticRoutes)))
	fmt.Printf("Avg bone: %s\n", time.Duration(totalBone)/time.Duration(len(bonehttprouter.StaticRoutes)))
	fmt.Printf("Avg httprouter: %s\n", time.Duration(totalHttprouter)/time.Duration(len(bonehttprouter.StaticRoutes)))
	fmt.Printf("Avg router: %s\n", time.Duration(totalRouter)/time.Duration(len(bonehttprouter.StaticRoutes)))
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

	fmt.Printf("%s \t Mean: %s \t P99: %s \t Max: %s\n", host+route.Path, metrics.Latencies.Mean, metrics.Latencies.P99, metrics.Latencies.Max)
	return metrics.Latencies.P99
}
