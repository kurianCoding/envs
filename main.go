package main

import "github.com/kurianCoding/envs/rnv"
import "fmt"

func main() {
	for key, value := range rnv.ReadEvns() {
		fmt.Println(key, value)
	}
	return
}
