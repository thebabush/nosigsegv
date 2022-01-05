package main

import (
	"fmt"

	"github.com/thebabush/nosigsegv"
)

func main() {
	errCode := nosigsegv.ErrorCode()

	switch errCode {
	case 1:
		fmt.Println("mmap failed :(")
	case 2:
		fmt.Println("mmap returned a non-fixed address :(")
	}
	if errCode != 0 {
		fmt.Println("you need to run `sudo sysctl -w vm.mmap_min_addr=0` before using this library!")
	} else {
		fmt.Println("initialization ok!")
	}

	fmt.Println("testing...")

	var ptr *uint64 = nil
	*ptr = 123
	fmt.Println("*", ptr, " = ", *ptr)
}
