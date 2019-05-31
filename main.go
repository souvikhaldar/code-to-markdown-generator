package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func init() {

}

func main() {
	var s, d, l string
	flag.StringVar(&s, "s", "", "full path to the source file")
	flag.StringVar(&d, "d", "", "full path to the destination file")
	flag.StringVar(&l, "l", "go", "name of the language, for eg. go,python,clojure")
	flag.Parse()
	fmt.Printf("Input source: %s, destination:%s and language: %s\n", s, d, l)
	inp, er := os.Open(s)
	if er != nil {
		log.Fatal("Error in opening the source file: ", er)
	}

	dest, er := os.Open(d)
	if er != nil {
		log.Println("Error in opening the destination file: ", er)
		log.Println("Creating the destination file")
		dest, err := os.Create(d)
		if err != nil {
			log.Fatalln("Error in creating the destination file: ", err)
		}
	}
	
}
