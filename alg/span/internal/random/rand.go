package random

import (
	"math/rand"
	"time"
)

var seed int64

func init() {
	seed = time.Now().UTC().UnixNano()
}

func Rand() *rand.Rand { _ = "STUB: not implemented"; return nil }
