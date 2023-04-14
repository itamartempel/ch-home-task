package valuedurliterator

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

// ValuedUrlIterator implement iterator that fetch data from buffer reader line by line
type ValuedUrlIterator struct {
	scanner *bufio.Scanner
}

func NewValuedUrlIterator(reader io.Reader) *ValuedUrlIterator {
	return &ValuedUrlIterator{
		scanner: bufio.NewScanner(reader),
	}
}

func (vui *ValuedUrlIterator) HasNextValuedUrl() bool {
	return vui.scanner.Scan()
}

func (vui *ValuedUrlIterator) GetNextValuedUrl() (url string, value int64) {
	line := vui.scanner.Text()

	// Extract the URL and integer from the line
	parts := strings.Split(line, " ")
	if len(parts) == 2 {
		url = parts[0]
		intStr := parts[1]

		// Convert the integer string to int64
		intVal, err := strconv.ParseInt(intStr, 10, 64)
		if err == nil {
			value = intVal
			return
		}
	}
	// since Im expected that the file will be currect i do not propagate the error
	log.Printf("Failed to parse valued url from line: %s", line)
	return
}
