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
		log.Fatalf("must use -c | --count flag and give an int")
	}

	n, err := strconv.Atoi(name)

	if err != nil {

		log.Printf("--count flag needs to be an integer: %s", err)
	}

	return n
}
