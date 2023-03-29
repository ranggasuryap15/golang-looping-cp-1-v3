package main

import "fmt"

func CountingNumber(n int) float64 {
	var total float64
	for i := 1; i < n; i++ {
		total += float64(i)
		total += float64(i) + 0.5
	}
	total += float64(n)
	return total // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(10))
}
