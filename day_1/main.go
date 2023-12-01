package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var nums = map[string]string{
	"1":"1","2":"2","3":"3","4":"4","5":"5","6":"6","7":"7","8":"8","9":"9",
	"one":"1","two":"2","three":"3","four":"4","five":"5","six":"6","seven":"7","eight":"8","nine":"9",
}

// map[string]int{
// 	"1":1,"2":2,"3":3,"4":4,"5":5,"6":6,"7":7,"8":8,"9":9,
// 	"one":1,"two":2,"three":3,"four":4,"five":5,"six":6,"seven":7,"eight":8,"nine":9,
// }

func main() {
	f, err := os.Open("day_1/input.txt")
	if err != nil{
		fmt.Println("file broke")
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	count := 0
	for scanner.Scan() {
		digits := ""
		line := scanner.Text()
		
		index:=0
		for index <= len(line) -1 {
			for numString, numChar := range nums{
				if strings.HasPrefix(line[index:], numString){
					digits = digits + numChar
				}
			}
			index++
		}
		
		digits = digits[:1] + digits[len(digits)-1:]

		valueInt, err:= strconv.Atoi(digits)
		if err != nil{
			log.Fatal()
		}
		count += valueInt 
	}
	

	println(count)
}