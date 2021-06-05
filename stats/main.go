package main

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

func main() {

	var d stats.Float64Data = []float64{1, 2, 3, 4, 4, 5}

	min, _ := d.Min()
	fmt.Println(min) // 1

	max, _ := d.Max()
	fmt.Println(max) // 5

	sum, _ := d.Sum()
	fmt.Println(sum) // 19

	mean, _ := d.Mean()
	fmt.Println(mean)

}
