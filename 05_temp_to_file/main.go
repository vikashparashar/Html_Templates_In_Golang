// creating a new html file store temp value in that file

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var firstName = os.Args[1]
	var lastName = os.Args[2]
	var myName = firstName + " " + lastName
	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<h1>My Name Is ` + strings.ToUpper(myName) + `</h1>
	</body>
	</html>
	`
	file, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	// io.Copy takes 2 arguments .
	// 1. First is dest
	// 2. Second argument is the io.Reader
	len, err := io.Copy(file, strings.NewReader(tpl))

	if err != nil {
		log.Fatalln("failed to copy data into recent created file")
	}

	fmt.Println("data successfully copied into recently created file . \n Copied Data Size : ", len)
}

// strings.ToUpper will convert the string text to capital text
