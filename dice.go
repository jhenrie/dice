package dice

import (
	"math/rand"
	"time"
)

type Dice struct {
	sides int
	r *rand.Rand
}

func NewDice(sides int) (*Dice) {
	source := rand.NewSource(time.Now().UnixNano())
	return &Dice {
		sides: sides,
		r: rand.New(source),
	}
}

func (me *Dice) Roll() int {
	return me.r.Intn(me.sides) + 1
}

func (me *Dice) RollN(count int) []int {
	results := make([]int, count)

	for i :=0; i < count; i++ {
		results[i] = me.Roll()
	}

	return results
}

