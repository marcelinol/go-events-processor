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

func LoadCompiledFile(h map[string]int) {
	finalPath := "./files/conversions"
	if _, err := os.Stat(finalPath); !os.IsNotExist(err) { // Check if file exists
		f, err := os.Open(finalPath)
		check(err)
		ProcessFile(f, h)
	}

	return
}

func readFiles() {
	path := "./tmp"
	files, err := ioutil.ReadDir(path)
	check(err)

	if len(files) < 1 {
		return
	}

	h := make(map[string]int)
	LoadCompiledFile(h)

	for _, file := range files {
		filePath := path + "/" + file.Name()
		file, err := os.Open(filePath)
		check(err)
		defer file.Close()

		ProcessFile(file, h)
		err = os.Remove(filePath)
		check(err)
	}

	WriteToFile(h)
}

func WriteToFile(hash map[string]int) {
	f, err := os.Create("./files/conversions")
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

func ProcessFile(f *os.File, hash map[string]int) {
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
	for range time.Tick(d) {
		fmt.Println("ping")
		f()
	}
}

func helloworld() {
	fmt.Printf("Hello, World!\n")
}

func main() {
	doEvery(7000*time.Millisecond, readFiles)
}
