package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
)

func main() {
	fmt.Print("$ ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	args := strings.Split(line," ")

	for i:=0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}
