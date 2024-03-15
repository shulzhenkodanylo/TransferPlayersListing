package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gosuri/uitable"
)

type Player struct {
	name          string
	team          string
	position      string
	transferValue int
	age           int
	nationality   string
}

func OpenPlayerList(pList []Player) {

	table := uitable.New()
	table.MaxColWidth = 50

	table.AddRow("INDEX", "|", "NAME", "|", "TEAM", "|", "POSITION", "|", "PRICE", "|", "AGE", "|", "NATIONALITY", "|")
	for i, p := range pList {
		table.AddRow(i+1, "|", p.name, "|", p.team, "|", p.position, "|", p.transferValue, "|", p.age, "|", p.nationality, "|")
	}
	fmt.Println(table)
}

func MakeTransfer(p Player, pIndex int, list []Player) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter a club name to which you want to move %v\n", p.name)
	scanner.Scan()
	transferTeam := scanner.Text()
	fmt.Printf("Enter a price for which %v will move to %v\n", p.name, transferTeam)
	scanner.Scan()
	transferPrice, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	list[pIndex].team = transferTeam
	list[pIndex].transferValue = transferPrice

}

func CorrentTransferStatus(pList []Player) string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter player index in list: ")

	scanner.Scan()
	playerNumberInList, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	for i, p := range pList {
		if i+1 == playerNumberInList {
			fmt.Printf("%v in %v\n", p.name, p.team)
			MakeTransfer(p, i, pList)
		}

	}
	return ("Player not found.")

}

func AddPlayer(pList *[]Player) {
	scanner := bufio.NewScanner(os.Stdin)
	newPlayer := Player{}
	for i := 0; i <= 5; i++ {
		switch i {
		case 0:
			fmt.Println("Enter this player NAME: ")
			scanner.Scan()
			newPlayer.name = scanner.Text()
		case 1:
			fmt.Println("Enter this player TEAM: ")
			scanner.Scan()
			newPlayer.team = scanner.Text()
		case 2:
			fmt.Println("Enter this player POSITION: ")
			scanner.Scan()
			newPlayer.position = scanner.Text()

		case 3:
			fmt.Println("Enter this player PRICE: ")
			scanner.Scan()
			scanTransferValue, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			newPlayer.transferValue = scanTransferValue
		case 4:
			fmt.Println("Enter this player AGE: ")
			scanner.Scan()
			scanAge, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			newPlayer.age = scanAge
		case 5:
			fmt.Println("Enter this player NAT: ")
			scanner.Scan()
			newPlayer.nationality = scanner.Text()
		}
	}
	*pList = append(*pList, newPlayer)
}

func UserFunctionality(pList []Player) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("[o]pen player list | [m]ake transfer | [a]dd player | [q]uit")
		scanner.Scan()
		userChoice := scanner.Text()

		switch strings.ToLower(userChoice) {
		case "o":
			OpenPlayerList(pList)
		case "m":
			CorrentTransferStatus(pList)
		case "a":
			AddPlayer(&pList)
		case "q":
			os.Exit(0)
		default:

		}
	}
}
func main() {
	playerList := []Player{
		{name: "Mikhaylo Mudryk", team: "Chelsea", position: "LW", transferValue: 42_000_000, age: 23, nationality: "Ukrainian"},
		{name: "Oleksandr Zinchenko", team: "Arsenal", position: "LD", transferValue: 42_000_000, age: 27, nationality: "Ukrainian"},
		{name: "Artem Dovbyk", team: "Girona", position: "ST", transferValue: 28_000_000, age: 26, nationality: "Ukrainian"},
		{name: "Vitaliy Mykolenko", team: "Everton", position: "LD", transferValue: 28_000_000, age: 24, nationality: "Ukrainian"},
		{name: "Ilya Zabarnyi", team: "Bournemoth", position: "CD", transferValue: 28_000_000, age: 21, nationality: "Ukrainian"},
		{name: "Viktor Tsygankov", team: "Girona", position: "RW", transferValue: 25_000_000, age: 26, nationality: "Ukrainian"},
		{name: "Anatoliy Trubin", team: "Benfica", position: "GK", transferValue: 25_000_000, age: 21, nationality: "Ukrainian"},
		{name: "Georgiy Sudakov", team: "Shakhtar", position: "AM", transferValue: 18_000_000, age: 21, nationality: "Ukrainian"},
		{name: "Andrii Lunin", team: "Real Madrid", position: "GK", transferValue: 8_000_000, age: 25, nationality: "Ukrainian"},
	}
	UserFunctionality(playerList)

}
