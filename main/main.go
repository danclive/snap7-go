package main

import (
	"fmt"

	"github.com/danclive/snap7-go"
)

func main() {
	snap7_client, err := snap7.ConnentTo("127.0.0.1", 0, 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(snap7_client)
}
