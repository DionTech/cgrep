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
	Mode       string
}

var files = make(chan string)

func (scan *Scan) find(path string, wg *sync.WaitGroup) {
	/** done this to tet concurrency
	fmt.Println(path)
	time.Sleep(500 * time.Millisecond)
	wg.Done()**/

	reg := regexp.MustCompile(scan.Expression)

	if scan.Mode == "find" {
		matches := reg.FindStringSubmatch(path)
		if len(matches) > 0 {
			fmt.Printf("%s:1: \n", path)
		}
		wg.Done()
		return
	}

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

func (scan *Scan) walk(files chan string, wg *sync.WaitGroup) {
	for path := range files {
		scan.find(path, wg)
	}
}

func (scan *Scan) Run() {
	wg := sync.WaitGroup{}

	for i := 0; i < scan.Threads; i++ {
		go scan.walk(files, &wg)
	}

	filepath.Walk(scan.Path, func(osPath string, f os.FileInfo, err error) error {
		f, err = os.Stat(osPath)

		// If no error
		if err != nil {
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
