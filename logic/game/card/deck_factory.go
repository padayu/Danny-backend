package card

import (
	"math/rand"
)

const (
	memoryCardLimit = 14
)

func GenerateMainDeck() []Memory {
	var deck = make([]Memory, 0)
	deck = append(deck, GenerateCards()...)

	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	deck = deck[:memoryCardLimit]

	return deck
}
