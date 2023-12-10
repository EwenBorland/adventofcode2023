package day_7

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func ParseSolution() string {
	f, err := os.Open("day_7/input.txt")
	if err != nil {
		fmt.Println("file broke")
		return "error"
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	partOneAnswer, partTwoAnswer := Day_7(scanner)

	return fmt.Sprintf("Part 1 Answer: %v\nPart 2 Answer: %v\n", partOneAnswer, partTwoAnswer)
}

func Day_7(scanner *bufio.Scanner) (int, int) {
	hands := []Hand{}
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		hands = append(hands, IdentifyHand(lineSplit[0], lineSplit[1]))
	}

	sort.Slice(hands, func(i, j int) bool {
		return !CompareHands(hands[i], hands[j])
	})

	totalWinnings := 0
	for i:=0; i < len(hands); i++{
		// fmt.Printf("hand: %v, bid: %v, rank: %v")
		totalWinnings += hands[i].bid * (i+1)
	}

	return 6440, totalWinnings
}

type Hand struct {
	cardsString string
	cardsInt    []int
	handType    int
	bid int
}

// handTypeMap := map[int]string{
// 	1:"Five of a kind",
// 	2:"Four of a kind",
// 	3:"Full house",
// 	4:"Three of a kind",
// 	5:"Two pair",
// 	6:"One pair",
// 	7:"High card",
// }

// CompareHands returns true if hand 'a' is higher value than hand 'b'
func CompareHands(a, b Hand) bool {
	comp := a.handType - b.handType
	if comp < 0 {
		return true
	} else if comp > 0 {
		return false
	} else {
		return CompareCards(a.cardsInt, b.cardsInt)
	}
}

// CompareCards compares the cards by index, returning true/false if one card is higher value
func CompareCards(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}
		return a[i] < b[i]
	}

	fmt.Printf("All cards evaluated as equal for sets a: %v, b:%v", a, b)
	return false
}

var cardMap = map[rune]int{
	'A': 0,
	'K': 1,
	'Q': 2,
	'J': 13,
	'T': 4,
	'9': 5,
	'8': 6,
	'7': 7,
	'6': 8,
	'5': 9,
	'4': 10,
	'3': 11,
	'2': 12,
}

func IdentifyHand(cards, bid string) Hand {
	h := Hand{cardsString: cards}

	cardsInt := []int{}
	for _, c := range cards {
		cardsInt = append(cardsInt, cardMap[c])
	}
	h.cardsInt = cardsInt

	h.handType = GetHandType(cardsInt)

	bidInt, err := strconv.Atoi(bid)
	if err!= nil{
		fmt.Println(err)
	}

	h.bid = bidInt

	return h
}

var handTypeSliceMap = map[int][]int{
	1: {5},
	2: {4, 1},
	3: {3, 2},
	4: {3, 1, 1},
	5: {2, 2, 1},
	6: {2, 1, 1, 1},
	7: {1, 1, 1, 1, 1},
}

func GetHandType(cards []int) int {
	cardMap := map[int]int{}

	for _, card := range cards {
		cardMap[card] = cardMap[card] + 1
	}

	jokers, ok := cardMap[13]
	if ok {
		delete(cardMap, 13)
	}

	thing := []int{}
	for _, v := range cardMap {
		thing = append(thing, v)
	}

	sort.Slice(thing, func(i, j int) bool {
		return thing[i] > thing[j]
	})

	fmt.Println(cards)

	if len(thing) > 0{
		thing[0] += jokers
	} else{
		thing = append(thing, jokers)
	}

	handType := 0
	for k, v := range handTypeSliceMap {
		// fmt.Printf("Does %v equal %v : ", thing, v)
		if slices.Equal(v, thing) {
			// fmt.Printf("Yes\n")
			handType = k
			break
		}

		// fmt.Printf("No\n")
	}

	return handType
}
