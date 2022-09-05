package mtga

import (
	"testing"
)

func TestParse(t *testing.T) {
	t.Run("simple parse card count", func(t *testing.T) {
		decklist, err := parseDeck(testDeck)

		AssertNil(t, err)

		totalCards := 0

		for _, card := range decklist {
			totalCards += card.Num
		}

		AssertEquals(t, totalCards, 75)
	})

	t.Run("set and collector number", func(t *testing.T) {
		decklist, err := parseDeck(testComplex)

		AssertNil(t, err)

		AssertEquals(t, decklist[0].Set, "WAR")
		AssertEquals(t, decklist[1].Set, "M21")

		AssertEquals(t, decklist[0].CollectorNumber, 9)
		AssertEquals(t, decklist[1].CollectorNumber, 15)
	})
}

func AssertNil(t *testing.T, got interface{}) {
	t.Helper()
	if got != nil {
		t.Fatalf(" got: %v\nwant: <nil>", got)
	}
}

func AssertEquals[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf(" got: %v\nwant: %v", got, want)
	}
}

const testDeck = `Deck
4 Beanstalk Giant
4 Bonecrusher Giant
4 Brazen Borrower
4 Breeding Pool
4 Edgewall Innkeeper
3 Escape to the Wilds
4 Fae of Wishes
6 Forest
3 Island
4 Lovestruck Beast
4 Lucky Clover
2 Mountain
2 Nissa, Who Shakes the World
3 Steam Vents
4 Stomping Ground
1 Temple of Abandon
2 Temple of Epiphany
2 Temple of Mystery
 
Sideboard
2 Aether Gust
1 Disdainful Stroke
1 Domri's Ambush
1 Escape to the Wilds
1 Expansion // Explosion
1 Fling
1 Grafdigger's Cage
1 Negate
1 Once and Future
1 Redcap Melee
1 Return to Nature
1 Shadowspear
1 Storm's Wrath
1 Unsummon`

const testComplex = `2 Defiant Strike (WAR) 9
	2 Defiant Strike (M21) 15`
