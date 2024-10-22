package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
)

func main() {
	for {
		fmt.Print("$ ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		trim := strings.TrimSpace(line)// trim the line of whitespace 
		args := strings.Split(trim," ")
		switch args[0] {
		case "mkdir":
			err = os.Mkdir(args[1], 0777)
			if err != nil {
				log.Fatal(err)
			}
		case "cd" :
			err = os.Chdir(args[1])
			if err != nil {
				log.Fatal(err)
			}
		case "pwd" :
			err = os.Chdir(args[1])
			if err != nil {
				log.Fatal(err)
			}
		case "hostname" :
			hostname, err := os.Hostname()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf(hostname)
		default:
			
			
		}
	}

}
