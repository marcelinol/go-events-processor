package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFiles() {
	h := make(map[string]int)
	files, err := ioutil.ReadDir("./tmp")
	check(err)

	for _, file := range files {
		fmt.Println(file.Name())
		file, err := os.Open("./tmp/" + file.Name())
		check(err)

		proccessFile(file, h)
		defer file.Close()
	}

	WriteToFile(h)
}

func WriteToFile(hash map[string]int) {
	f, err := os.Create("./files/conversions" + strconv.Itoa(int(time.Now().UnixNano())))
	check(err)
	defer f.Close()

	var buffer bytes.Buffer

	for key, value := range hash {
		buffer.WriteString(key + ":" + strconv.Itoa(value) + "\n")
	}

	_, err = f.WriteString(buffer.String())
	check(err)

	return
}

func proccessFile(f *os.File, hash map[string]int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() { // O(n) sendo n o numero de linhas do arquivo
		fmt.Println(scanner.Text())
		values := strings.Split(scanner.Text(), ":") // O(m) sendo m o numero de caracteres da linha
		email := values[0]
		count, err := strconv.Atoi(values[1])
		check(err)

		_, emailExist := hash[email]
		if emailExist {
			hash[email] = hash[email] + count
		} else {
			hash[email] = count
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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
	doEvery(5000*time.Millisecond, readFiles)
}
