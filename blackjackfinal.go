package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var picks = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
var colors = [4]string{"spade", "hearts", "plum", "diamonds"}
var cardValue = map[string]int{"A": 11, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "J": 10, "Q": 10, "K": 10}

//initialize the card by connecting picks and colors to one deck. Create a deck of cards with 52 pieces by using a slice.
func initCard() []string {
	var cards = make([]string, 0)
	for _, color := range colors {
		for _, pick := range picks {
			cards = append(cards, color+"_"+pick)
		}
	}
	return cards[:52]
}

//sort the slice cards randomly in order to shuffle the deck and then generate a new out-of-order cards.
func shuffleCard(cards []string) []string {
	rr := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := len(cards)
	for i := l - 1; i > 0; i-- {
		r := rr.Intn(i)
		cards[r], cards[i] = cards[i], cards[r]
	}
	return cards
}

//the function for dealer decision
func dealer_decision_func(dealer []string, player []string) int {

	return_val := 2

	if count(dealer) < 17 {
		//fmt.Println("condition1")
		return_val = 1
	} else if count(dealer) >= 21 {
		//fmt.Println("condition2")
		return_val = 2
	} else if count(dealer) < count(player) && count(player) <= 21 {
		//fmt.Println("condition3")
		return_val = 1
	}

	return return_val
}

//the function for calculate the value of cards you take(both for dealer and user)
func count(cards []string) int {
	picks := make([]string, 1)
	hasA := 0
	for _, card := range cards {
		card = strings.Split(card, "_")[1]
		picks = append(picks, card)
		if card == "A" {
			hasA += 1
		}
	}
	count := 0
	for _, card := range picks {
		count += cardValue[card]
	}
	for true {
		if hasA > 0 {
			if count > 21 {
				count -= 10
				hasA -= 1
			} else {
				break
			}
		} else {
			break
		}
	}
	return count
}

//generate the global variable to scan
var inputReader *bufio.Reader
var err error
var input string
var choice string

//main function for judgement
func main() {

	cards := initCard()
	cards = shuffleCard(cards)

	var player = make([]string, 2)
	var dealer = make([]string, 2)
	player = []string{cards[1], cards[3]}
	dealer = []string{cards[0], cards[2]}

	fapai := 4

	user_decision := 1
	dealer_decision := 1

	fmt.Print("Do you want to play? yes or no")
	inputReader = bufio.NewReader(os.Stdin)
	choice, err = inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	} else if choice == "yes\n" {

		//"beat the dealer situation 1" Blackjack
		if count(dealer) == 21 {
			user_decision = 0
			dealer_decision = 0
			fmt.Println("You lose. Dealer got BlackJack.", dealer)
		} else if count(player) == 21 {
			user_decision = 0
			dealer_decision = 0
			fmt.Println("You win! You got BlackJack.", player)
		}
		//decision of user and dealing the cards to continue the game
		for user_decision == 1 || dealer_decision == 1 {

			if user_decision == 1 {
				fmt.Println("------------ New Round ------------")
				fmt.Println("dealer's cards: ", dealer, "your cards: ", player)

				if len(player) == 2 {
					fmt.Println("One more card? 1. yes; 2. no; 3. surrender.")
				} else {
					fmt.Println("One more card? 1. yes; 2. no.")
				}

				reenter := 1

				for reenter == 1 {

					inputReader = bufio.NewReader(os.Stdin)
					input, err = inputReader.ReadString('\n')
					if err != nil {
						fmt.Println("There were errors reading, exiting program.")
						return
					}
					switch input {
					case "1\n":
						fmt.Println("You will get one more card.")
						user_decision = 1
						reenter = 0
					case "2\n":
						fmt.Println("Keep your current cards.")
						user_decision = 2
						reenter = 0
					case "3\n":
						if len(player) == 2 {
							user_decision = 3
							reenter = 0
						} else {
							fmt.Printf("Read error. Please re-enter.")
						}
					default:
						fmt.Printf("Read error. Please re-enter.")
					}
				}

			}

			if user_decision == 1 {
				player = append(player, cards[fapai])
				fapai += 1
			}

			// generate dealer's decision automatically
			if dealer_decision == 1 {

				dealer_decision = dealer_decision_func(dealer, player)

			}

			if dealer_decision == 1 {

				dealer = append(dealer, cards[fapai])
				fapai += 1

			}

		}

		//final result judgement for the game
		if user_decision != 0 && dealer_decision != 0 {
			fmt.Println("************* Game Result *************")
			if user_decision == 3 {
				fmt.Println("You surrender.")
			} else if count(player) <= 21 && count(dealer) < count(player) {
				fmt.Println("dealer's cards: ", dealer, "your cards: ", player)
				fmt.Println("You win!")
			} else if count(dealer) > 21 && count(player) <= 21 {
				fmt.Println("dealer's cards: ", dealer, "your cards: ", player)
				fmt.Println("You win!")
			} else {
				fmt.Println("dealer's cards: ", dealer, "your cards: ", player)
				fmt.Println("You didn't win.")
			}

		}
	} else if choice == "no\n" {
		os.Exit(3)
	}

}
