package strategies

import (
	"github.com/rodrigo-brito/ninjabot/strategy"
)

func NewStrategy() strategy.Strategy {
	return new(CrossEMA)
}
