package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <fileName>")
		return
	}

	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	countUnique := 0
	i := 0
	for scanner.Scan() {
		leftDiv := i % 1000
		if leftDiv == 0 {
			fmt.Println("Left next lines behind: ", i)
		}
		ip := scanner.Text()

		file2, err2 := os.Open(fileName)
		if err2 != nil {
			fmt.Println("Error opening file 2:", err2)
			return
		}

		scanner2 := bufio.NewScanner(file2)

		isFind := false
		for scanner2.Scan() {
			ip2 := scanner2.Text()
			if ip == ip2 {
				isFind = true
				break
			}
		}

		if err := scanner2.Err(); err != nil {
			fmt.Println("Error reading file2:", err)
			return
		}

		file2.Close()

		if !isFind {
			countUnique++
			fmt.Println("Found uniqueu address: ", countUnique)
		}

		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Found uniqueu address: ", countUnique)
}
