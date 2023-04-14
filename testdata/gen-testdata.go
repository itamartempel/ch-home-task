package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	CreateTestDataByFileSize(1024, "1MB.txt")
	CreateTestDataByFileSize(1024*10, "10MB.txt")
	CreateTestDataByFileSize(1024*100, "100MB.txt")
	// CreateTestDataByFileSize(1024*1024, "1GB.txt")
}

func CreateTestDataByFileSize(fileSizeKB int64, fileName string) {
	file, err := os.OpenFile(fmt.Sprintf("./testdata/generated/%s", fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Could not open", fileName)
		return
	}
	defer file.Close()
	var i = 0
	for {

		_, err2 := file.WriteString(fmt.Sprintf("http://api.tech.com/item/%d %d\n", 100_000+rand.Intn(900_000), rand.Intn(1000_000)))
		if err2 != nil {
			fmt.Println("Could not write text to ", fileName)
		}
		if i%1000 == 0 {
			file.Sync()
			stat, _ := file.Stat()
			if stat.Size()/1024 > fileSizeKB {
				break
			}
		}
		i++
	}

	fmt.Println(fileName, " Created successfully")
}
