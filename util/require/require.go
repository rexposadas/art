package require

import (
	"log"
	"strconv"
)

func FileName(name string) {
	if name == "" {
		log.Fatal("needs to specific -f | --file flag for file")
	}
}

func Count(name string) int {
	if name == "" {
		return 1
	}

	n, err := strconv.Atoi(name)
	if err != nil {
		log.Printf("--count flag needs to be an integer: %s", err)
	}

	if n < 1 {
		log.Printf("--count flag needs to be greater than 0")
	}

	return n
}
