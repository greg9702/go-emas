package randomizer

import (
	"errors"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// IRandomizer interface for radnomizer components
type IRandomizer interface {
	RandInt(min int, max int) (int, error)
}

var singleton *BaseRandomizer
var once sync.Once

// BaseRand static function returns randomizer
func BaseRand() *BaseRandomizer {
	once.Do(func() {
		singleton = newBaseRandomizer(time.Now().UnixNano())
	})
	return singleton
}

// BaseRandomizer is a base implementaion of IRandomizer
type BaseRandomizer struct {
}

func newBaseRandomizer(seed int64) *BaseRandomizer {
	b := BaseRandomizer{}
	rand.Seed(seed)
	return &b
}

// RandInt returns random value from [min, max] interval, error while passed min > max or negative values
func (b *BaseRandomizer) RandInt(min int, max int) (int, error) {
	if min > max {
		return 0, errors.New("Passed wrong values min: " + strconv.Itoa(min) + " max: " + strconv.Itoa(max))
	}
	return rand.Intn(max-min+1) + min, nil
}
