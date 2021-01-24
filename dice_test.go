package dice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDice( t *testing.T) {
	count := 20
	d6 := NewDice(6)
	d20 := NewDice(20)
	t.Run("dice roll in bounds for d6", func(t *testing.T) {
		for i := 0; i < count; i++ {
			result := d6.Roll()
			assert.True(t, result > 0 && result <= 6, fmt.Sprintf("D6 Bounds Dice Test: %d - Result: %d", i, result))
		}
	})

	t.Run("dice roll in bounds for d20", func(t *testing.T){
		for i := 0; i < count; i++ {
			result := d20.Roll()
			assert.True(t, result > 0 && result <= 20, fmt.Sprintf("D20 Bounds Dice Test: %d - Result: %d", i, result))
		}
	})

	t.Run("sequence roll of d6 all in bounds", func(t *testing.T){
		results := d6.RollN(count)
		assert.Len(t, results, count, "lenght test of series result")

		for i := 0; i < count; i++ {
			assert.True(t, results[i] > 0 && results[i] <= 6, fmt.Sprintf("D6 Bounds Dice Test: %d - Result: %d", i, results[i]))
		}
	})

	t.Run("sequence roll of d20 all in bounds", func(t *testing.T){
		results := d20.RollN(count)
		assert.Len(t, results, count, "lenght test of series result")

		for i := 0; i < count; i++ {
			assert.True(t, results[i] > 0 && results[i] <= 20, fmt.Sprintf("D20 Bounds Dice Test: %d - Result: %d", i, results[i]))
		}
	})
}
