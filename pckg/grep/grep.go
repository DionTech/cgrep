package grep

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

type Scan struct {
	Path       string
	Expression string
	Threads    int
}

var files = make(chan string)

func Find(path string, expression string, wg *sync.WaitGroup) {
	/** done this to text concurrency
	fmt.Println(path)
	time.Sleep(500 * time.Millisecond)
	wg.Done()**/

	reg := regexp.MustCompile(expression)
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}

	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		matches := reg.FindStringSubmatch(scanner.Text())
		if len(matches) > 0 {
			for range matches {
				fmt.Printf("%s:%d: %s\n", path, i, scanner.Text())
			}
		}
	}

	wg.Done()
}

func Walk(files chan string, expression string, wg *sync.WaitGroup) {
	for path := range files {
		Find(path, expression, wg)
	}
}

func (scan *Scan) Run() {
	//files := make(chan string, scan.Threads)

	wg := sync.WaitGroup{}

	for i := 0; i < scan.Threads; i++ {
		go Walk(files, scan.Expression, &wg)
	}

	filepath.Walk(scan.Path, func(osPath string, f os.FileInfo, err error) error {
		f, err = os.Stat(osPath)

		// If no error
		if err != nil {
			//fmt.Println(err)

			return nil
		}

		// File & Folder Mode
		fMode := f.Mode()

		// Is folder
		if fMode.IsDir() {

		} else {
			wg.Add(1)
			go func(osPath string) {
				files <- osPath
			}(osPath)
		}
		return nil
	})

	wg.Wait()
}
