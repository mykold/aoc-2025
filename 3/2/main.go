package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func intify(b []byte) (r int64) {
	m := int64(1)
	for i := len(b) - 1; i >= 0; i-- {
		r += int64(b[i]-'0') * m
		m *= 10
	}
	return
}

const k = 12

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var result int64
	for scanner.Scan() {
		line := scanner.Text()

		pop := len(line) - k

		stack := make([]byte, 1, len(line))
		stack[0] = line[0]
		for i := 1; i < len(line); i++ {
			d := line[i]

			for len(stack) > 0 && pop > 0 && d > stack[len(stack)-1] {
				stack = stack[:len(stack)-1] // pop
				pop--
			}

			stack = append(stack, d) // push
		}

		result += intify(stack[:k])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
