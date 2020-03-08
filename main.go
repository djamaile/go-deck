package main

import "fmt"

func main(){
	cards := newDeck()
	cards.print()
	cards.shuffle()
	fmt.Println()
	fmt.Println()
	cards.print()
	fmt.Println()
	fmt.Println()
	cards.print()
	fmt.Println()
	fmt.Println()
	cards.print()
}
