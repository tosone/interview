package main

import (
	"fmt"
	"unsafe"
)

// S ..
type S struct {
	E struct {
		x bool
		y int8
	}
	A uint32
	B uint64
	C uint64
	D uint64
}

func main() {
	fmt.Printf("A offset: %d\n", unsafe.Offsetof(S{}.A))
	fmt.Printf("E offset: %d\n", unsafe.Offsetof(S{}.E))
	fmt.Printf("E size: %d\n", unsafe.Sizeof(S{}.E))
	fmt.Printf("Total size: %d\n", unsafe.Sizeof(S{}))
}
