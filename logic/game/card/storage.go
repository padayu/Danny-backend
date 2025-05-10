package card

import "strconv"

func GenerateCards() []Memory {
	var fullMemoryDeck = make([]Memory, 60)
	for index := range fullMemoryDeck {
		fullMemoryDeck[index] = Memory{Id: index, Image: strconv.Itoa(index) + ".png"}
	}
	return fullMemoryDeck
}
