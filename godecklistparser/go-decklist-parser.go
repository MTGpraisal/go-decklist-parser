package godecklistparser

func Parse(s string) ([]Card, error) {
	return parseMTGA(s)
}

// Card represents a parsed decklist card
type Card struct {
	Num             int
	Name            string
	Set             string
	CollectorNumber int
}

func newCard(num int, name, set string, collectorNumber int) Card {
	return Card{
		Num:             num,
		Name:            name,
		Set:             set,
		CollectorNumber: collectorNumber,
	}
}
