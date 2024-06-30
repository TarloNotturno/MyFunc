package MyFunc

import (
	"math/rand/v2"
	"time"
)

func GetNumbRnd() int {
	app1 := time.Now()
	app := app1.Nanosecond()
	s3 := rand.NewPCG(0, (uint64)(app))
	r3 := rand.New(s3)
	return r3.IntN(100)
}
