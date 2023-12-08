package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func day01() {
	readFile, err := os.Open("inputs/1.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	result := int64(0)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		re, _ := regexp.Compile(`(\d)`)
		digitsGroups := re.FindAllStringSubmatch(line, -1)

		first, _ := strconv.ParseInt(digitsGroups[0][0], 10, 32)
		last, _ := strconv.ParseInt(digitsGroups[len(digitsGroups)-1][0], 10, 32)
		result += first*10 + last
	}

	fmt.Println("result:", result)

	readFile.Close()
}
