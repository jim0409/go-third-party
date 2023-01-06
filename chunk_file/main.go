package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

func main() {
	split := 4

	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	texts := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		texts = append(texts, text)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	lengthPerSplit := len(texts) / split
	for i := 0; i < split; i++ {
		if i+1 == split {
			chunkTexts := texts[i*lengthPerSplit:]
			writefile(strings.Join(chunkTexts, "\n"))
		} else {
			chunkTexts := texts[i*lengthPerSplit : (i+1)*lengthPerSplit]
			writefile(strings.Join(chunkTexts, "\n"))
		}
	}
}

func writefile(data string) {
	file, err := os.Create("chunks-" + uuid.New().String() + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.WriteString(data)
}
