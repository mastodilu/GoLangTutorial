package main

import (
	txa "./textanalyze"
)

func main() {
	txa.CountWords(txa.Quote)
	txa.PrintMap()
}
