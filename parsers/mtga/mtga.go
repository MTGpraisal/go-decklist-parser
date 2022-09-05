package mtga

import (
	"errors"
	"strconv"
	"strings"

	godecklistparser "github.com/MTGpraisal/go-decklist-parser"
)

var exclusions = []string{
	" ",
	"deck",
	"decklist",
	"library",
	"sideboard",
	"commander",
	"companion",
}

var (
	ErrNotACard            = errors.New("not a card")
	ErrInvalidCollectorNum = errors.New("collector number not an int")
)

func parseDeck(decklist string) ([]godecklistparser.Card, error) {
	cards := strings.Split(decklist, "\n")
	parsedCards := make([]godecklistparser.Card, 0, len(cards))

	for i := range cards {
		parsedCard, err := parseCard(cards[i])

		switch err {
		case nil:
			parsedCards = append(parsedCards, parsedCard)

		case ErrNotACard:
			continue

		case ErrInvalidCollectorNum:

			return nil, err
		}
	}

	return parsedCards, nil
}

func newCard(num int, name, set string, collectorNumber int) godecklistparser.Card {
	return godecklistparser.Card{
		Num:             num,
		Name:            name,
		Set:             set,
		CollectorNumber: collectorNumber,
	}
}

// parseCard is a bodge, but it works well enough for my needs
func parseCard(card string) (godecklistparser.Card, error) {
	var ok = false
	num := 1
	set := ""
	collectorNumber := 0

	tokenised := strings.Split(card, " ")

	// One word card (ie shock), or a structural word
	if len(tokenised) <= 1 {
		if len(tokenised) == 0 || contains(exclusions, strings.ToLower(tokenised[0])) {
			return godecklistparser.Card{}, ErrNotACard
		} else {
			return newCard(num, tokenised[0], "", 0), nil
		}
	}

	// strings.Split(" ", " ") returns {"", ""}
	if tokenised[0] == "" && tokenised[1] == "" {
		return godecklistparser.Card{}, ErrNotACard
	}

	// Have at least 2 things to look at, so we start parsing proper
	num, tokenised = parseNum(tokenised)

	if len(tokenised) == 1 {
		return newCard(num, tokenised[0], "", 0), nil
	}

	// Assuming card can't be in format of "(SET) num", so we assume there can't be a collector number
	if len(tokenised) == 2 {
		set, tokenised = parseSet(tokenised)
		return newCard(num, strings.Join(tokenised, " "), set, 0), nil
	}

	// Check the last & then second to last tokens
	set, ok = isSetCode(tokenised[len(tokenised)-1])
	if ok {
		return newCard(num, strings.Join(tokenised[:len(tokenised)-1], " "), set, 0), nil
	}

	set, ok = isSetCode(tokenised[len(tokenised)-2])
	if !ok {
		return newCard(num, strings.Join(tokenised, " "), set, collectorNumber), nil
	}

	collectorNumber, err := strconv.Atoi(tokenised[len(tokenised)-1])
	if err != nil {
		return godecklistparser.Card{}, ErrInvalidCollectorNum
	}
	return newCard(num, strings.Join(tokenised[:len(tokenised)-2], " "), set, collectorNumber), nil

}

func parseNum(card []string) (int, []string) {
	num, err := strconv.Atoi(card[0])
	if err == nil {
		return num, card[1:]
	} else {
		return 1, card
	}
}

func parseSet(card []string) (string, []string) {
	set, ok := isSetCode(card[len(card)-1])
	if ok {
		return set, card[0 : len(card)-2]
	} else {
		return "", card
	}
}

func isSetCode(s string) (string, bool) {
	if len(s) == 5 && strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") {
		return s[1:4], true
	}

	return "", false
}

func contains(xs []string, s string) bool {
	for _, x := range xs {
		if s == x {
			return true
		}
	}

	return false
}
