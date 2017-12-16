package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoves(t *testing.T) {
	//- Spin, written sX, makes X programs move from the end to the front, but maintain their order otherwise. (For example, s3 on abcde produces cdeab).
	//- Exchange, written xA/B, makes the programs at positions A and B swap places.
	//- Partner, written pA/B, makes the programs named A and B swap places.
	t.Run("Spin", func(t *testing.T) {
		state, l, index := initialState()
		// a b c ... k l m n o p
		// v
		// v
		// v
		// n o p a b c ... k l m
		spin(state, l, index, 3)
		assert.Equal(t, 'n', state[0])
		assert.Equal(t, 'o', state[1])
		assert.Equal(t, 'p', state[2])
		assert.Equal(t, 'a', state[3])
		assert.Equal(t, 'b', state[4])
		assert.Equal(t, 'c', state[5])
		assert.Equal(t, 'k', state[l-3])
		assert.Equal(t, 'l', state[l-2])
		assert.Equal(t, 'm', state[l-1])

		assert.Equal(t, 0, index['n'])
		assert.Equal(t, 1, index['o'])
		assert.Equal(t, 2, index['p'])
		assert.Equal(t, 3, index['a'])
		assert.Equal(t, 4, index['b'])
		assert.Equal(t, 5, index['c'])
		assert.Equal(t, l-3, index['k'])
		assert.Equal(t, l-2, index['l'])
		assert.Equal(t, l-1, index['m'])
	})
	t.Run("Exchange", func(t *testing.T) {
		state, l, index := initialState()
		exchange(state, l, index, 0, 1)
		assert.Equal(t, 'b', state[0])
		assert.Equal(t, 'a', state[1])
		assert.Equal(t, 1, index['a'])
		assert.Equal(t, 0, index['b'])
	})
	t.Run("Partner", func(t *testing.T) {
		state, l, index := initialState()
		partner(state, l, index, 'a', 'b')
		assert.Equal(t, 'b', state[0])
		assert.Equal(t, 'a', state[1])
		assert.Equal(t, 1, index['a'])
		assert.Equal(t, 0, index['b'])
	})
}
