package main

import (
	"fmt"

	"github.com/danclive/snap7-go"
)

// func main() {
// 	snap7_client, err := snap7.ConnentTo("127.0.0.1", 0, 0)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(snap7_client)
// }

func main() {
	host := "116.85.47.167"
	client, err := snap7.ConnentTo2(host, 0x0200, 0x0200, 0)

	fmt.Println(client)
	fmt.Println(err)
}
