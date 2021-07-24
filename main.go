package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}

	cards.print()

}

func newCard() string {
	return "five of diamonds"
}
