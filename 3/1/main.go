package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var result int32
	xs := make([]int32, 0, 100)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		for i := range len(line) {
			for j := range i {
				xs = append(
					xs,
					(rune(line[i])-'0')+((rune(line[j])-'0')*10),
				)
			}
		}

		var r int32
		for _, x := range xs {
			if x > r {
				r = x
			}
		}

		result += r
		xs = xs[:0]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
