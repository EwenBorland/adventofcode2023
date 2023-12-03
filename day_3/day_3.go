package day_3

import (
	"bufio"
	"fmt"
	"os"
)

func ParseSolution() string {
	f, err := os.Open("day_3/input.txt")
	if err != nil{
		fmt.Println("file broke")
		return "error"
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	validCount := Day_3(scanner)
	
	return fmt.Sprintf("Valid Sum: %v\n", validCount, )
}


/*

idea 1: 
masking maybe

initialise the 0th and 1st line with falses

if a symbol is parsed then set the appropriate indexes in the mask true

apply the mask to the parts somehow

idea 2:
the symbols only affect parts that are on adjacent lines, parse line n for parts, then parse lines n-1,n,and n+1 for symbols and match indices 
*/

func Day_3(scanner *bufio.Scanner) (int) {
	setup := false
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if !setup{
			setup = true
			
		}

		
		

	}

	return 0

}