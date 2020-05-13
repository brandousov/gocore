// App Core functions and vars | Math
// @help:
// fmt.Printf("LeadingZeros64(%064b) = %d\n", 1, bits.LeadingZeros64(1))
// https://golang.org/pkg/math/bits/#LeadingZeros64
package core

import (
	"math/rand"
	"sync/atomic"
)





//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////// |structs|
// Инкрементальный счётчик для многопоточных задач
func AtomicAdd(sum *uint64, delta uint64) uint64 {
	return atomic.AddUint64(sum, delta)
}





// Random
func RandInt64() int64 {
	//rand.Seed(time.Now().UTC().UnixNano())
	//return rand.Int63n(math.MaxInt64)
	return rand.Int63()
}





// Random from min to max
func Rand(min, max int64) int64 {
	return rand.Int63n(max - min) + min
}
