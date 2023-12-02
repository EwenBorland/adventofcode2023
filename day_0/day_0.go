package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day_0/input.txt")
	if err != nil{
		fmt.Println("file broke")
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		for _, value := range strings.Split(line, ","){
			valueInt, err:= strconv.Atoi(value)
			if err != nil{
				fmt.Println("atoi broke")
				return
			}
			count += valueInt
		}        
	}

	println(count)
}
