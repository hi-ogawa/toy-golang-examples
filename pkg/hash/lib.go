package hash

import (
	"unsafe"
)

// By Chris Wellons https://nullprogram.com/blog/2018/07/31/
func Hash11u(x uint32) uint32 {
	x ^= x >> 16
	x *= 0x7feb352d
	x ^= x >> 15
	x *= 0x846ca68b
	x ^= x >> 16
	return x
}

func Hash21u(x [2]uint32) uint32 {
	return Hash11u(Hash11u(x[0]) + x[1])
}

func Hash31u(x [3]uint32) uint32 {
	return Hash11u(Hash11u(Hash11u(x[0])+x[1]) + x[2])
}

func Hash11(x float32) float32 {
	return u2f(Hash11u(f2u(x)))
}

func Hash21(x [2]float32) float32 {
	v := [2]uint32{f2u(x[0]), f2u(x[1])}
	return u2f(Hash21u(v))
}

func Hash23(x [2]float32) [3]float32 {
	v1 := [3]uint32{f2u(x[0]), f2u(x[1]), 1}
	v2 := [3]uint32{f2u(x[0]), f2u(x[1]), 2}
	return [3]float32{Hash21(x), u2f(Hash31u(v1)), u2f(Hash31u(v2))}
}

// They are not inverse each other
func u2f(x uint32) float32 {
	return float32(x) / (1 << 32)
}

func f2u(x float32) uint32 {
	return *(*uint32)(unsafe.Pointer(&x))
}
