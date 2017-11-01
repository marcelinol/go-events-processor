package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile() {
	files, err := ioutil.ReadDir("./tmp")
	check(err)

	for _, file := range files {
		fmt.Println(file.Name())
		file, err := os.Open("./tmp/" + file.Name())
		check(err)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

}

func doEvery(d time.Duration, f func()) {
	count := 0
	for range time.Tick(d) {
		f()
		count++
		if count > 0 {
			return
		}
	}
}

func helloworld() {
	fmt.Printf("Hello, World!\n")
}

func main() {
	doEvery(2000*time.Millisecond, helloworld)
	doEvery(5000*time.Millisecond, readFile)
}
