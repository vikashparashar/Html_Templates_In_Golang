// taking arguments as value for myName Variable

package main

import (
	"fmt"
	"os"
)

func main() {
	var myName = os.Args[1]
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
		<h1>My Name Is ` + myName + `</h1>
	</body>
	</html>
	`

	fmt.Println(tpl)
}
