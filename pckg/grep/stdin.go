package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sync"
)

var texts = make(chan string)

func GrepText(reg *regexp.Regexp, wg *sync.WaitGroup) {

	for text := range texts {
		matches := reg.FindStringSubmatch(text)
		if len(matches) > 0 {
			for range matches {
				fmt.Printf("%s\n", text)
			}
		}
		wg.Done()
	}

}

func Grep(expression string, threads int) {
	reg := regexp.MustCompile(expression)

	scanner := bufio.NewScanner(os.Stdin)
	wg := sync.WaitGroup{}

	for i := 0; i < threads; i++ {
		go GrepText(reg, &wg)
	}

	for {
		// Scans a line from Stdin(Console)
		scanner.Scan()

		// Holds the string that scanned
		text := scanner.Text()

		if len(text) != 0 {
			wg.Add(1)

			go func() {
				texts <- text
			}()

		} else {
			break
		}
	}

	wg.Wait()
}
