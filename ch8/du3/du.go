package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range directoryEntries(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			n.Add(1)
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// sema limits the total amount of open "files"
var sema = make(chan struct{}, 20)

func directoryEntries(dir string) []os.FileInfo {
	sema <- struct{}{}        // aquire tokens
	defer func() { <-sema }() // release tokens when we are done
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	var n sync.WaitGroup

	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files,  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
