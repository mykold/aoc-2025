// --- Part Two ---
// The clerk quickly discovers that there are still invalid IDs in the ranges in your list. Maybe the young Elf was doing other silly patterns as well?
//
// Now, an ID is invalid if it is made only of some sequence of digits repeated at least twice. So, 12341234 (1234 two times), 123123123 (123 three times), 1212121212 (12 five times), and 1111111 (1 seven times) are all invalid IDs.
//
// From the same example as before:
//
// 11-22 still has two invalid IDs, 11 and 22.
// 95-115 now has two invalid IDs, 99 and 111.
// 998-1012 now has two invalid IDs, 999 and 1010.
// 1188511880-1188511890 still has one invalid ID, 1188511885.
// 222220-222224 still has one invalid ID, 222222.
// 1698522-1698528 still contains no invalid IDs.
// 446443-446449 still has one invalid ID, 446446.
// 38593856-38593862 still has one invalid ID, 38593859.
// 565653-565659 now has one invalid ID, 565656.
// 824824821-824824827 now has one invalid ID, 824824824.
// 2121212118-2121212124 now has one invalid ID, 2121212121.
// Adding up all the invalid IDs in this example produces 4174379265.
//
// What do you get if you add up all of the invalid IDs using these new rules?

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	input := string(b)

	xs := make(map[int]struct{})

	for _, r := range strings.Split(input, ",") {
		p := strings.Split(r, "-")

		fs, ls := p[0], p[1]

		fi, err := strconv.Atoi(fs)
		if err != nil {
			log.Fatal(err)
		}
		li, err := strconv.Atoi(ls)
		if err != nil {
			log.Fatal(err)
		}

		for i := fi; i <= li; i++ {
			s := strconv.Itoa(i)

			for j := 1; j <= len(s)/2; j++ {
				var skip bool

				t := s[:j]

				if len(s)%len(t) != 0 {
					continue
				}

				for k := 0; k < len(s); k += len(t) {
					if t == s[k:k+len(t)] {
						continue
					}
					skip = true
				}

				if !skip {
					xs[i] = struct{}{}
				}
			}

		}
	}

	var result int
	for k := range xs {
		result += k
	}
	fmt.Println(result)
}
