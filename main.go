package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var countUnique int
var mu sync.Mutex

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <fileName> <numProc>")
		return
	}

	fileName := os.Args[1]
	numProc, err := strconv.Atoi(os.Args[2])
	if err != nil || numProc == 0 {
		fmt.Println("You forget set param <numProc>")
		return
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	searchIPChan := make(chan string, numProc)

	for c := 0; c < numProc; c++ {
		go scanIP2(searchIPChan, fileName)
	}

	// countUnique := 0
	i := 0
	for scanner.Scan() {
		leftDiv := i % 1000
		if leftDiv == 0 {
			fmt.Println("Left next lines behind: ", i)
		}
		ip := scanner.Text()

		searchIPChan <- ip

		i++
	}

	close(searchIPChan)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Found uniqueu address: ", countUnique)
}

func scanIP2(searchIPChan <-chan string, fileName string) {

	fmt.Println("Started proc")

	for ip := range searchIPChan {
		file2, err2 := os.Open(fileName)
		if err2 != nil {
			fmt.Println("Error opening file 2:", err2)
			fmt.Println("IP will need to check egain:", ip)
			continue
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

		file2.Close()

		if err := scanner2.Err(); err != nil {
			fmt.Println("Error reading file2:", err)
		}

		if !isFind {
			mu.Lock()
			countUnique++
			fmt.Println("Found uniqueu address: ", countUnique)
			mu.Unlock()
		}
	}
}
