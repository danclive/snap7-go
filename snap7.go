package main

import (
	"fmt"
)

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L./ -lsnap7
#include "snap7.h"
*/
import "C"

func main() {
	a := C.errIsoConnect

	fmt.Println(a)

	client := C.Cli_Create()

	fmt.Println(client)
}
