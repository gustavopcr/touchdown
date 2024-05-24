package main

import (
	"fmt"

	"github.com/gustavopcr/touchdown/api"
)

func main() {
	err := api.NewServer(8080)
	if err != nil {
		fmt.Println("error: ", err)
	}
}
