// // a simlple html template with custom data

package main

import "fmt"

func main() {
	var myName string = "Vikash Parashar"
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
