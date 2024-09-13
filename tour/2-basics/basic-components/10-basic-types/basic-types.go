package main

import (
	"fmt"
	"math/cmplx"
)

var (
	Boolean bool       = false
	MaxInt  uint64     = 1<<64 - 1
	Complex complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", Boolean, Boolean)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", Complex, Complex)
}

/*

Basic types in Go:

string

bool

int		int8	int16		int32		int64
uint	uint8	uint16	uint32	uint64	uintptr

A uint is an unsigned (non-negative) integer

byte	-	alias for uint8

rune	-	alias for int32
			- represents a Unicode code point

float32		float64
complex64	complex128

*/
