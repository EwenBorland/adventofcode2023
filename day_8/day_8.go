package day_8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ParseSolution() string {
	f, err := os.Open("day_8/input.txt")
	if err != nil {
		fmt.Println("file broke")
		return "error"
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	partOneAnswer, partTwoAnswer := Day_8(scanner)

	return fmt.Sprintf("Part 1 Answer: %v\nPart 2 Answer: %v\n", partOneAnswer, partTwoAnswer)
}

func Day_8(scanner *bufio.Scanner) (int, int) {
	scanner.Scan()
	instructions := scanner.Text()

	nodes := map[string]Node{}
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		n := ParseNode(line)
		nodes[n.name] = n
		// fmt.Println(n)
	}

	totalInstructions := 0
	// currentNode := nodes["AAA"]

	// atZZZ := false
	// for !atZZZ {
	// 	for _, v := range instructions {
	// 		// fmt.Printf("Current Node: %v, instructions: %v\n",currentNode.name, totalInstructions)
	// 		if currentNode.name == "ZZZ" {
	// 			atZZZ = true
	// 			break
	// 		}
	// 		totalInstructions++
	// 		if v == 'L' {
	// 			currentNode = nodes[currentNode.left]
	// 		} else {
	// 			currentNode = nodes[currentNode.right]
	// 		}
	// 	}
	// }

	//part 2

	currentNodes := []Node{}
	for _, v := range nodes {
		if strings.HasSuffix(v.name, "A") {
			currentNodes = append(currentNodes, v)
		}
	}
	fmt.Println("nodes: ",currentNodes)

	totalInstructions2 := 0
	allAtZ := false
	completePathLengths := []int{}
	for !allAtZ {
		for _, v := range instructions {
			fmt.Printf("Current Node: %v, instructions: %v\n",currentNodes[0].name, totalInstructions2)
			deletables := []int{}
			for i, n := range currentNodes {
				if strings.HasSuffix(n.name, "Z") {
					fmt.Printf("Node: %v, has suffix\n", n)
					completePathLengths = append(completePathLengths, totalInstructions2)
					deletables = append(deletables, i)
				}
				fmt.Printf("Node: %v, deos not have suffix\n", n)
			}

			fmt.Printf("Deletables: %v, lengths: %v\n",deletables, completePathLengths)
			
			for i, j := range deletables {
				currentNodes = append(currentNodes[:(j-i)], currentNodes[(j-i)+1:]...)
			}

			if len(currentNodes) == 0 {
				allAtZ = true
				break
			}

			totalInstructions2++
			if v == 'L' {
				for i, n := range currentNodes {
					currentNodes[i] = nodes[n.left]
				}
			} else {
				for i, n := range currentNodes {
					currentNodes[i] = nodes[n.right]
				}
			}
		}
	}

	fmt.Println("aaa?",completePathLengths)
	a:=LCM(completePathLengths[0], completePathLengths[1], completePathLengths[2:]...)
	// this LCM func that I copied from some random place on the interwebs didn't work for the full solution
	// had to manually take the complete path lengths and use some other random website to calculate the correct answer.


	return totalInstructions, a
}

type Node struct {
	s     string
	name  string
	left  string
	right string
}

func ParseNode(line string) Node {

	lineSplit := strings.Split(line, " = ")
	node := Node{s: line, name: lineSplit[0]}

	line2 := strings.ReplaceAll(lineSplit[1], " ", "")
	line2 = strings.ReplaceAll(line2, "(", "")
	line2 = strings.ReplaceAll(line2, ")", "")
	lineSplit = strings.Split(line2, ",")

	node.left = lineSplit[0]
	node.right = lineSplit[1]

	return node
}

func GCD(a, b int) int {
	for b != 0 {
			t := b
			b = a % b
			a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	if len(integers) == 0{
		return result
	}

	for i := 0; i < len(integers); i++ {
			result = LCM(result, integers[i])
	}

	return result
}