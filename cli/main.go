package main

import (
	"bufio"
	"ch-home-task/pkg/topnvaluedurls"
	"ch-home-task/pkg/valuedurliterator"
	"fmt"
	"log"
	"os"
)

func main() {

	filePath := GetFilePathFromStdin()
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to open provedied file", err)
	}
	defer file.Close()

	// Create iterator that will read the file line by line and deserialize it
	iterator := valuedurliterator.NewValuedUrlIterator(file)
	// Create a priority queue to keep track of the top 10 URLs with the biggest integers
	topTenUrls := topnvaluedurls.NewTopNValuedUrls(10)

	// Stitch both components
	for iterator.HasNextValuedUrl() {
		url, value := iterator.GetNextValuedUrl()
		topTenUrls.AddValuedUrl(url, value)
	}

	// Print the results
	for _, url := range topTenUrls.GetTopNUrls() {
		fmt.Println(url)
	}
}

func GetFilePathFromStdin() string {
	stdinScanner := bufio.NewScanner(os.Stdin)
	stdinScanner.Scan()
	return stdinScanner.Text()
}
