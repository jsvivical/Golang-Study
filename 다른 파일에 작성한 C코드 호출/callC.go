package main

//#cgo CFLAGS : -I${SRCDIR}/callClib
//#cgo LDFLAGS : ${SRCDIR}/callC.a
//#include <stdlib.h>
//#include "callC.h"

import "C"
import "fmt"

func main() {
	fmt.Println("going to call a C function\n")
	C.cHello()
}
