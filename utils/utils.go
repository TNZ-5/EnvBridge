package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ListAllVars() {
	env := os.Environ()
	for _, envVar := range env {
		fmt.Println(envVar)
	}
}

func SearchVar(target string) {
	env := os.Environ()

	for _, envVar := range env {

		pair := strings.SplitN(envVar, "=", 2)

		key := pair[0]
		value := pair[1]

		if key == target {
			fmt.Printf("Found: %v %v\n", key, value)
			return
		}

	}

	fmt.Println("Not Found")

}

func SetFile(filePath string){
	file, err := os.Open(filePath)

	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()

	

	reader := bufio.NewReader(file)
	
	for{
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		parts := strings.SplitN(line,"=",2)
		os.Setenv(parts[0],parts[1])
	}

}

func SaveToSession(key string, value string){
	os.Setenv(key,value)
}

func SaveToFile(key string, value string){
	file, err := os.OpenFile("files/default.env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("failed to open file: %s", err)
    }
    defer file.Close()

    _, err = file.WriteString(fmt.Sprintf("%s=%s\n", key, value))
    if err != nil {
        log.Fatalf("failed to write to file: %s", err)
    }
}