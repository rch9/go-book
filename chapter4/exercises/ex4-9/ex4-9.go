package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	counts := make(map[string]int)
	wordfreq("file", counts)
	fmt.Println(counts)
}

func wordfreq(filename string, counts map[string]int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
