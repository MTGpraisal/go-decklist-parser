# go-decklist-parser

This is a package to parse MTGA format decklists into the following format
```
type Card struct {
	Num             int
	Name            string
	Set             string
	CollectorNumber int
}
```

This package is very much WIP and not guaranteed to be reliable until release v1.0

Other decklist formats will be added as demand dictates, or when found necessary by the primary MTGpraisal project.