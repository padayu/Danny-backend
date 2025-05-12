package game

import "math/rand"

func GenerateWordOptions() []string {
	wordOptions := make([][]string, 0)
	wordOptions = append(wordOptions, []string{"sas", "sus", "sos"})
	wordOptions = append(wordOptions, []string{"banana", "bobr", "baba yaga"})
	wordOptions = append(wordOptions, []string{"lol", "kek", "cheburek"})
	wordOptions = append(wordOptions, []string{"shaurma", "shaverma", "shavuha"})
	randomIndex := rand.Intn(len(wordOptions))
	return wordOptions[randomIndex]
}
