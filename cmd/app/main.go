package main

import (
	"fmt"
	"github.com/muesli/clusters"
	"github.com/muesli/kmeans"
	"math/rand"
)

func main() {
	for x := 0; x < 4; x++ {
		go makeCluster()
	}
	makeCluster()
}

func makeCluster() {

	var d clusters.Observations
	for x := 0; x < 500000; x++ {
		d = append(d, clusters.Coordinates{
			rand.Float64(),
			rand.Float64(),
		})
	}
	//	Partition the data points into 16 clusters
	km := kmeans.New()
	clusters, _ := km.Partition(d, 16)

	for _, c := range clusters {
		//	fmt.Printf(" %.2f ", c.Center[0])
		fmt.Print(len(c.Observations), " ")

		//	fmt.Printf("Centered at x: %.2f y: %.2f\n", c.Center[0], c.Center[1])
		//	fmt.Printf("Matching data points: %+v\n\n", c.Observations)
	}
}
